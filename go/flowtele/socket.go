package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/scionproto/scion/go/flowtele/dbus/datalogger"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/flowtele"
	"github.com/scionproto/scion/go/flowtele/dbus"
	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/log"
	sd "github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"
	"github.com/scionproto/scion/go/lib/snet/squic"
	"github.com/scionproto/scion/go/lib/sock/reliable"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	errorNoError quic.ErrorCode = 0x100
)

var (
	localIAFromFlag, remoteIAFromFlag addr.IA
	scionPath                         scionPathDescription

	remoteIpFlag       = kingpin.Flag("ip", "IP address to connect to").Default("127.0.0.1").String()
	remotePortFlag     = kingpin.Flag("port", "Port number to connect to").Default("5500").Int()
	useRemotePortRange = kingpin.Flag("port-range", "Use increasing (remote) port numbers for additional QUIC senders").Default("false").Bool()
	localIpFlag        = kingpin.Flag("local-ip", "IP address to listen on (required for SCION)").Default("").String()
	localPortFlag      = kingpin.Flag("local-port", "Port number to listen on (required for SCION)").Default("0").Int()
	useLocalPortRange  = kingpin.Flag("local-port-range", "Use increasing local port numbers for additional QUIC senders").Default("true").Bool()
	quicSenderOnly     = kingpin.Flag("quic-sender-only", "Only start the quic sender").Default("false").Bool()
	fshaperOnly        = kingpin.Flag("fshaper-only", "Only start the fshaper").Default("false").Bool()
	quicDbusIndex      = kingpin.Flag("quic-dbus-index", "index of the quic sender dbus name").Default("0").Int()
	nConnections       = kingpin.Flag("num", "Number of QUIC connections").Default("2").Int()
	noApplyControl     = kingpin.Flag("no-apply-control", "Do not forward apply-control calls from fshaper to this QUIC connection (useful to ensure the calibrator flow is not influenced by vAlloc)").Default("false").Bool()
	mode               = kingpin.Flag("mode", "the sockets mode of operation: fetch, quic, fshaper").Default("fetch").String()
	maxData            = kingpin.Flag("max-data", "the maximum amount of data that should be transmitted on each QUIC flow (0 means no limit)").Default("0").Int()

	useScion        = kingpin.Flag("scion", "Open scion quic sockets").Default("false").Bool()
	dispatcherFlag  = kingpin.Flag("dispatcher", "Path to dispatcher socket").Default("").String()
	sciondAddrFlag  = kingpin.Flag("sciond", "SCIOND address").Default(sd.DefaultAPIAddress).String()
	scionPathsFile  = kingpin.Flag("paths-file", "File containing a list of SCION paths to the destination").Default("").String()
	scionPathsIndex = kingpin.Flag("paths-index", "Index of the path to use in the --paths-file").Default("0").Int()

	rate          = kingpin.Flag("rate", "Fixed rate in Mbits/s").Default("0").Uint64()
	csvFilePrefix = kingpin.Flag("csv-prefix", "File prefix to use for writing the CSV file.").Default("rtt").String()
	localIAFlag   = kingpin.Flag("local-ia", "ISD-AS address to listen on.").String()
	remoteIAFlag  = kingpin.Flag("remote-ia", "ISD-AS address to connect to.").String()
	scionPathFlag = kingpin.Flag("path", "SCION path to use.").String()
)

var (
	sigs = make(chan os.Signal, 1)
	done = make(chan bool, 1)
)

const (
	Bit  = 1
	KBit = 1000 * Bit
	MBit = 1000 * KBit
	GBit = 1000 * MBit

	Byte  = 8 * Bit
	KByte = 1000 * Byte
	MByte = 1000 * KByte
)

func setAddrIA(s string) (t *addr.IA) {
	t = &addr.IA{}
	if s == "" {
		return
	}

	if err := t.Set(s); err != nil {
		log.Debug(fmt.Sprintf("Not setting addrIA to %v\n", s))
	}
	return
}

func setScionPath(s string) (t *scionPathDescription) {
	t = &scionPathDescription{}
	if s == "" {
		return
	}

	if err := t.Set(s); err != nil {
		log.Debug(fmt.Sprintf("Not setting scionPath to %v\n", s))
		log.Error(fmt.Sprintf("Setting scionPath: %v\n", err))
	}
	return
}

func init() {
	setupLogger()
	kingpin.Parse()
	localIAFromFlag = *setAddrIA(*localIAFlag)
	log.Debug(fmt.Sprintf("LocalIA %v\n", localIAFromFlag))
	remoteIAFromFlag = *setAddrIA(*remoteIAFlag)
	scionPath = *setScionPath(*scionPathFlag)
}

func main() {
	// first run
	// python3.6 athena_m2.py 2
	// clear; go run go/flowtele/quic_listener.go --num 3
	// clear; go run go/flowtele/socket.go --fshaper-only
	// clear; go run go/flowtele/socket.go --quic-sender-only --ip 164.90.176.95 --port 5500 --quic-dbus-index 0
	// clear; go run go/flowtele/socket.go --quic-sender-only --ip 164.90.176.95 --port 5501 --quic-dbus-index 1
	// clear; go run go/flowtele/socket.go --quic-sender-only --ip 164.90.176.95 --port 5502 --quic-dbus-index 2
	// can add --no-apply-control to calibrator flow

	// ./scion.sh topology -c topology/Tiny.topo
	// ./scion.sh start
	// bazel build //... && bazel-bin/go/flowtele/listener/linux_amd64_stripped/flowtele_listener --scion --sciond 127.0.0.12:30255 --local-ia 1-ff00:0:110 --num 2
	// bazel build //... && bazel-bin/go/flowtele/linux_amd64_stripped/flowtele_socket --quic-sender-only --scion --sciond 127.0.0.19:30255 --local-ip 127.0.0.1 --local-port 6000 --ip 127.0.0.1 --port 5500 --local-ia 1-ff00:0:111 --remote-ia 1-ff00:0:110 --path 1-ff00:0:111,1-ff00:0:110
	// bazel build //... && bazel-bin/go/flowtele/linux_amd64_stripped/flowtele_socket --quic-sender-only --scion --sciond 127.0.0.19:30255 --local-ip 127.0.0.1 --local-port 6001 --ip 127.0.0.1 --port 5501 --local-ia 1-ff00:0:111 --remote-ia 1-ff00:0:110 --path 1-ff00:0:111,1-ff00:0:110
	errChannel := make(chan error)
	closeChannel := make(chan struct{})

	log.Info("Starting...")
	if *quicSenderOnly || *mode == "quic" {
		log.Info("QUIC sender\n")
		invokeQuicSenders(closeChannel, errChannel)
	} else if *fshaperOnly || *mode == "fshaper" {
		log.Info("FShaper\n")
		go func(cc chan struct{}, ec chan error) {
			defer log.HandlePanic()
			invokeFshaper(cc, ec)
		}(closeChannel, errChannel)
	} else if *mode == "fetch" {
		log.Info("Fetch\n")
		go func(cc chan struct{}, ec chan error) {
			defer log.HandlePanic()
			invokePathFetching(cc, ec)
		}(closeChannel, errChannel)
	} else {
		flag.PrintDefaults()
		errChannel <- fmt.Errorf("Must provide either --quic-sender-only or --fshaper-only")
	}

	select {
	case err := <-errChannel:
		log.Error(fmt.Sprintf("Error encountered (%s), exiting socket\n", err))
		os.Exit(1)
	case <-closeChannel:
		log.Info("Exiting without errors")
	}
}

func setupLogger() {
	logCfg := log.Config{Console: log.ConsoleConfig{Level: "debug"}}
	if err := log.Setup(logCfg); err != nil {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "Error configuring logger. Exiting due to:%s\n", err)
		os.Exit(-1)
	}
}

func invokePathFetching(closeChannel chan struct{}, errChannel chan error) {
	sciondAddr := *sciondAddrFlag
	localIA := localIAFromFlag
	remoteIA := remoteIAFromFlag
	paths, err := fetchPaths(sciondAddr, localIA, remoteIA)
	if err != nil {
		errChannel <- err
	} else {
		for _, path := range paths {
			fmt.Println(NewScionPathDescription(path).String())
		}
		close(closeChannel)
	}
}

func invokeQuicSenders(closeChannel chan struct{}, errChannel chan error) {
	// start QUIC instances
	// TODO(cyrill) read flow specs from config/user_X.json
	log.Info(fmt.Sprintf("Starting %d QUIC senders:\n", *nConnections))
	remoteIp := net.ParseIP(*remoteIpFlag)
	localIp := net.ParseIP(*localIpFlag)
	var wg sync.WaitGroup
	for i := 0; i < *nConnections; i++ {
		wg.Add(1)
		go func(index int) {
			defer log.HandlePanic()
			defer wg.Done()
			localPort := *localPortFlag
			remotePort := *remotePortFlag

			if *useLocalPortRange {
				localPort += index
			}
			if *useRemotePortRange {
				remotePort += index
			}

			localAddr := net.UDPAddr{IP: localIp, Port: localPort}
			remoteAddr := net.UDPAddr{IP: remoteIp, Port: remotePort}
			err := startQuicSender(&localAddr, &remoteAddr, int32(*quicDbusIndex+index), !*noApplyControl, errChannel)
			if err != nil {
				errChannel <- err
			}
		}(i)
	}
	go func() {
		defer log.HandlePanic()
		wg.Wait()
		close(closeChannel)
	}()
}

func invokeFshaper(closeChannel chan struct{}, errChannel chan error) {
	fdbus := flowteledbus.NewFshaperDbus(*nConnections)

	// if a min interval for the fshaper is specified, make sure to accumulate acked bytes that would otherwise not be registered by athena
	// fdbus.SetMinIntervalForAllSignals(5 * time.Millisecond)

	// dbus setup
	if err := fdbus.OpenSessionBus(); err != nil {
		errChannel <- err
		return
	}

	// register method and listeners
	if err := fdbus.Register(); err != nil {
		errChannel <- err
		return
	}

	// listen for feedback from QUIC instances and forward to athena
	go func() {
		defer log.HandlePanic()
		for v := range fdbus.SignalListener {
			if fdbus.Conn.Names()[0] == v.Sender {
				// fdbus.Log("ignore signal %s generated by socket", v.Name)
			} else {
				// fdbus.Log("forwarding signal...")
				signal := flowteledbus.CreateFshaperDbusSignal(v)
				if fdbus.ShouldSendSignal(signal) {
					if err := fdbus.Send(signal); err != nil {
						errChannel <- err
						return
					}
				}
			}
		}
	}()

	// don't close the closeChannel to keep running forever
	// fdbus.Close()
	// close(closeChannel)
}

func fetchPaths(sciondAddr string, localIA addr.IA, remoteIA addr.IA) ([]snet.Path, error) {
	sdConn, err := sd.NewService(sciondAddr).Connect(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize SCION network: %s", err)
	}
	if localIA.Equal(addr.IA{}) {
		log.Debug("fetchPaths: Got empty localIA. Fetching from SCIOND now...")
		localIA, err = sdConn.LocalIA(context.Background())
		if err != nil {
			log.Error(fmt.Sprintf("Error fetching localIA from SCIOND: %v\n", err))
		}
	}

	log.Debug(fmt.Sprintf("Remote: %v Local: %v\n", remoteIA, localIA))
	paths, err := sdConn.Paths(context.Background(), remoteIA, localIA, sd.PathReqFlags{})
	if err != nil {
		return nil, fmt.Errorf("Failed to lookup paths: %s", err)
	}
	return paths, nil
}

func fetchPath(pathDescription *scionPathDescription, sciondAddr string, localIA addr.IA, remoteIA addr.IA) (snet.Path, error) {
	paths, err := fetchPaths(sciondAddr, localIA, remoteIA)
	if err != nil {
		return nil, err
	}
	for _, path := range paths {
		if pathDescription.IsEqual(NewScionPathDescription(path)) {
			return path, nil
		}
	}
	return nil, fmt.Errorf("No matching path (%v) was found in %v", pathDescription, paths)
}

func establishQuicSession(localAddr *net.UDPAddr, remoteAddr *net.UDPAddr, tlsConfig *tls.Config, quicConfig *quic.Config) (quic.Session, error) {
	if *useScion {
		dispatcher := *dispatcherFlag
		sciondAddr := *sciondAddrFlag
		var pathDescription *scionPathDescription
		if !scionPath.IsEmpty() {
			pathDescription = &scionPath
		} else if *scionPathsFile != "" {
			pathDescriptions, err := readPaths(*scionPathsFile)
			if err != nil {
				return nil, fmt.Errorf("Couldn't read paths from file %s: %s", *scionPathsFile, err)
			}
			if *scionPathsIndex >= len(pathDescriptions) {
				return nil, fmt.Errorf("SCION path index out of range %d >= %d", *scionPathsIndex, len(pathDescriptions))
			}
			pathDescription = pathDescriptions[*scionPathsIndex]
		} else {
			return nil, fmt.Errorf("Must specify either --path or --paths-file and --paths-index")
		}
		localIA := localIAFromFlag
		remoteIA := remoteIAFromFlag

		// fetch path fitting to description
		var remoteScionAddr snet.UDPAddr
		remoteScionAddr.Host = remoteAddr
		remoteScionAddr.IA = remoteIA
		if !remoteIA.Equal(localIA) {
			path, err := fetchPath(pathDescription, sciondAddr, localIA, remoteIA)
			if err != nil {
				return nil, err
			}
			remoteScionAddr.Path = path.Path()
			remoteScionAddr.NextHop = path.UnderlayNextHop()
		}

		// setup SCION connection
		ds := reliable.NewDispatcher(dispatcher)
		sciondConn, err := sd.NewService(sciondAddr).Connect(context.Background())
		if err != nil {
			return nil, fmt.Errorf("Unable to initialize SCION network: %s", err)
		}
		network := snet.NewNetwork(localIA, ds, sd.RevHandler{Connector: sciondConn})

		// start QUIC session
		return squic.Dial(network, localAddr, &remoteScionAddr, addr.SvcNone, quicConfig)
	} else {
		// open UDP connection
		// localAddr := net.UDPAddr{IP: net.IPv4zero, Port: 0}
		conn, err := net.ListenUDP("udp", localAddr)
		if err != nil {
			fmt.Printf("Error starting UDP listener: %s\n", err)
			return nil, err
		}

		// start QUIC session
		return quic.Dial(conn, remoteAddr, "host:0", tlsConfig, quicConfig)
	}
}

func startQuicSender(localAddr *net.UDPAddr, remoteAddr *net.UDPAddr, flowId int32, applyControl bool, errChannel chan error) error {
	// capture interrupts to gracefully terminate run
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer log.HandlePanic()
		sig := <-sigs
		fmt.Printf("%v\n", sig)
		close(done)
	}()

	// setup datalogger
	dLogger := datalogger.NewDbusDataLogger(
		fmt.Sprintf("%s-samples-%d.csv", *csvFilePrefix, time.Now().Unix()), []string{"flowID", "microTimestamp", "microSRTT"},
		[]string{"src", "dest"},
	)
	defer dLogger.Close()
	dLogger.SetMetadata([]string{localAddr.String(), remoteAddr.String()})
	dLogger.Run()

	// start dbus
	qdbus := flowteledbus.NewQuicDbus(flowId, applyControl)
	qdbus.SetMinIntervalForAllSignals(5 * time.Millisecond)
	if err := qdbus.OpenSessionBus(); err != nil {
		return err
	}
	defer qdbus.Close()

	if err := qdbus.Register(); err != nil {
		return err
	}

	// signal forwarding functions
	newSrttMeasurement := func(t time.Time, srtt time.Duration) {
		if qdbus.Conn == nil {
			// ignore signals if the session bus is not connected
			return
		}

		if srtt > math.MaxUint32 {
			panic("srtt does not fit in uint32")
		}
		signal := flowteledbus.CreateQuicDbusSignalRtt(flowId, t, uint32(srtt.Microseconds()))
		dLogger.Send(&datalogger.RTTData{FlowID: int(flowId), Timestamp: t, SRtt: srtt})
		if qdbus.ShouldSendSignal(signal) {
			if err := qdbus.Send(signal); err != nil {
				fmt.Printf("srtt -> %d\n", qdbus.FlowId)
				errChannel <- err
			}
		}
	}

	packetsLost := func(t time.Time, newSlowStartThreshold uint64) {
		if qdbus.Conn == nil {
			// ignore signals if the session bus is not connected
			return
		}

		if newSlowStartThreshold > math.MaxUint32 {
			panic("newSlotStartThreshold does not fit in uint32")
		}
		signal := flowteledbus.CreateQuicDbusSignalLost(flowId, t, uint32(newSlowStartThreshold))
		if qdbus.ShouldSendSignal(signal) {
			if err := qdbus.Send(signal); err != nil {
				fmt.Printf("lost -> %d\n", qdbus.FlowId)
				errChannel <- err
			}
		}
	}
	packetsAcked := func(t time.Time, congestionWindow uint64, packetsInFlight uint64, ackedBytes uint64) {
		if qdbus.Conn == nil {
			// ignore signals if the session bus is not connected
			return
		}

		if congestionWindow > math.MaxUint32 {
			panic("congestionWindow does not fit in uint32")
		}
		if packetsInFlight > math.MaxInt32 {
			panic("packetsInFlight does not fit in int32")
		}
		if ackedBytes > math.MaxUint32 {
			panic("ackedBytes does not fit in uint32")
		}
		ackedBytesSum := qdbus.Acked(uint32(ackedBytes))
		signal := flowteledbus.CreateQuicDbusSignalCwnd(flowId, t, uint32(congestionWindow), int32(packetsInFlight), ackedBytesSum)
		if qdbus.ShouldSendSignal(signal) {
			if err := qdbus.Send(signal); err != nil {
				fmt.Printf("ack -> %d\n", qdbus.FlowId)
				errChannel <- err
			}
			qdbus.ResetAcked()
		}
	}

	flowteleSignalInterface := flowtele.CreateFlowteleSignalInterface(newSrttMeasurement, packetsLost, packetsAcked)
	// make QUIC idle timout long to allow a delay between starting the listeners and the senders
	quicConfig := &quic.Config{MaxIdleTimeout: time.Hour,
		FlowTeleSignal: flowteleSignalInterface}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, NextProtos: []string{"Flowtele"}}

	// setup quic session
	session, err := establishQuicSession(localAddr, remoteAddr, tlsConfig, quicConfig)
	if err != nil {
		return fmt.Errorf("Error starting QUIC connection to [%s]: %s", remoteAddr.String(), err)
	}
	defer func() {
		fmt.Printf("closing session %d\n", flowId)
		// session.Close()
		qdbus.Session = nil
	}()
	qdbus.Session = checkFlowTeleSession(session)

	// open stream
	//rateInBitsPerSecond := uint64(20 * 1000 * 1000)
	//session.SetFixedRate(rateInBitsPerSecond)
	//qdbus.Log("set fixed rate %f...", float64(rateInBitsPerSecond)/1000000)
	qdbus.Log("session established. Opening stream...")
	ctx, cancelStream := context.WithCancel(context.Background())
	go func() {
		defer log.HandlePanic()
	loop:
		for {
			select {
			case <-done:
				cancelStream()
				if err := session.CloseWithError(errorNoError, "Interrupt received."); err != nil {
					fmt.Printf("Error closing session: %v\n", err)
				}
				break loop
			}
		}
	}()
	stream, err := session.OpenStreamSync(ctx)
	if err != nil {
		return fmt.Errorf("Error opening QUIC stream to [%s]: %s", remoteAddr.String(), err)
	}
	defer func() {
		fmt.Println("closing stream")
		if err := stream.Close(); err != nil {
			fmt.Printf("Deffered Error closing stream: %v\n", err)
		}
	}()
	qdbus.Log("stream opened %d", stream.StreamID())
	// continuously send 10MB messages to quic listener
	message := make([]byte, 10000000)
	for i := range message {
		message[i] = 42
	}

	if *rate > 0 {
		qdbus.Log("Setting rate to %d Mbit/s", *rate)
		checkFlowTeleSession(qdbus.Session).SetFixedRate(*rate * MBit)
	}

	sentBytes := 0

loop:
	for {
		select {
		case <-done:
			break loop
		default:
			if *maxData == 0 {
				_, err = stream.Write(message)
			} else {
				if sentBytes < *maxData {
					var n int
					n, err = stream.Write(message[0:min(len(message), *maxData-sentBytes)])
					sentBytes += n
				} else {
					break
				}
			}
			if err != nil {
				return fmt.Errorf("Error writing message to [%s]: %s", remoteAddr.String(), err)
			}
		}
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkFlowTeleSession(s quic.Session) quic.FlowTeleSession {
	fs, ok := s.(quic.FlowTeleSession)
	if !ok {
		panic("Returned session is not flowtele sessions")
	}
	return fs
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}
