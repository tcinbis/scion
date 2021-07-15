package main

import (
	"crypto/tls"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/flowtele"
	"github.com/lucas-clemente/quic-go/http3"
	"github.com/netsec-ethz/scion-apps/pkg/shttp"
	slog "github.com/scionproto/scion/go/lib/log"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	dataDir = "go/src/scion/go/flowtele/stream/data/"
	Bit     = 1
	KBit    = 1000 * Bit
	MBit    = 1000 * KBit
	GBit    = 1000 * MBit

	Byte  = 8 * Bit
	KByte = 1000 * Byte
	MByte = 1000 * KByte
)

var (
	ip       = kingpin.Flag("ip", "ip to listen on").Default("127.0.0.1").String()
	port     = kingpin.Flag("port", "port the server listens on").Default("8001").Uint()
	useScion = kingpin.Flag("scion", "Enable serving server via SCION").Default("false").Bool()
)

func init() {
	kingpin.Parse()
	fmt.Printf("Parsing complete. Listening on %v:%v\n", *ip, *port)
}

func getTCPConn(addr string) *net.TCPListener {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	tcpConn, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	return tcpConn
}

func getUDPConn(addr string) *net.UDPConn {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal()
	}
	return udpConn
}

func getQuicConf() *quic.Config {
	newSrttMeasurement := func(t time.Time, srtt time.Duration) {
		//fmt.Printf("t: %v, srtt: %v\n", t, srtt)
	}

	flowteleSignalInterface := flowtele.CreateFlowteleSignalInterface(newSrttMeasurement, func(t time.Time, newSlowStartThreshold uint64) {

	}, func(t time.Time, lostRatio float64) {

	},
		func(t time.Time, congestionWindow uint64, packetsInFlight uint64, ackedBytes uint64) {

		})
	// make QUIC idle timout long to allow a delay between starting the listeners and the senders
	return &quic.Config{
		//MaxIdleTimeout: time.Second,
		//KeepAlive: true,
		FlowTeleSignal: flowteleSignalInterface,
	}
}

func startTCPServer(handler http.Handler) {
	fmt.Printf("Using QUIC\n")
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	certFile := "/home/tom/go/src/scion/go/flowtele/tls.pem"
	keyFile := "/home/tom/go/src/scion/go/flowtele/tls.key"

	quicServer := http3.Server{
		Server: &http.Server{
			Addr: addr,
		},
		QuicConfig: getQuicConf(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		quicServer.SetQuicHeaders(w.Header())
		handler.ServeHTTP(w, r)
	})
	mux.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		quicServer.SetQuicHeaders(w.Header())
		// Small 40x40 png
		w.Write([]byte{
			0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
			0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x28,
			0x01, 0x03, 0x00, 0x00, 0x00, 0xb6, 0x30, 0x2a, 0x2e, 0x00, 0x00, 0x00,
			0x03, 0x50, 0x4c, 0x54, 0x45, 0x5a, 0xc3, 0x5a, 0xad, 0x38, 0xaa, 0xdb,
			0x00, 0x00, 0x00, 0x0b, 0x49, 0x44, 0x41, 0x54, 0x78, 0x01, 0x63, 0x18,
			0x61, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x01, 0xe2, 0xb8, 0x75, 0x22, 0x00,
			0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
		})
	})
	//quicServer.Handler = mux
	quicServer.Server.Handler = mux

	certs := make([]tls.Certificate, 1)
	var err error
	certs[0], err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates: certs,
	}
	go func() {
		defer slog.HandlePanic()
		quicServer.Server.Serve(tls.NewListener(getTCPConn(addr), tlsConfig))
	}()
	go func() {
		defer slog.HandlePanic()
		quicServer.Serve(getUDPConn(addr))
	}()

	for {
		//for key, _ := range *quicServer.GetSessions() {
		//	fmt.Printf("ConnectionID: %s\n", (*key).ConnectionId())
		//}
		time.Sleep(1 * time.Second)
	}

	//http.Handle("/", handler)
	//http.ListenAndServe(fmt.Sprintf("%s:%d", *ip, *port), nil)
}

func startSCIONServer(handler http.Handler) {
	fmt.Printf("Using SCION\n")

	server := shttp.NewScionServer(fmt.Sprintf("%s:%d", *ip, *port), withLogger(handler), nil, getQuicConf())
	server.Server.SetNewStreamCallback(func(sess *quic.EarlySession, strID quic.StreamID) {
		fmt.Printf("%v %v\n", time.Now(), sess)
		fmt.Printf("%v: Session to %s open.\n", time.Now(), (*sess).RemoteAddr())
		fmt.Printf("%v %v\n", time.Now(), server.Server.GetSessions())
		(*checkFlowTeleSession(checkSession(sess))).SetFixedRate(500 * KByte)
	})

	log.Fatal(server.ListenAndServe())
}

func main() {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting user home dir.")
	}
	filePath := filepath.Join(userHome, dataDir, "bbb_30fps")
	fmt.Println(filePath)
	handler := http.FileServer(http.Dir(filePath))
	if *useScion {
		startSCIONServer(handler)
	} else {
		startTCPServer(handler)
	}

}

// withLogger returns a handler that logs requests (after completion) in a simple format:
//	  <time> <remote address> "<request>" <status code> <size of reply>
func withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrec := &recordingResponseWriter{ResponseWriter: w}
		h.ServeHTTP(wrec, r)

		log.Printf("%s \"%s %s %s/SCION\" %d %d\n",
			r.RemoteAddr,
			r.Method, r.URL, r.Proto,
			wrec.status, wrec.bytes)
	})
}

type recordingResponseWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (r *recordingResponseWriter) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *recordingResponseWriter) Write(b []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	r.bytes += len(b)
	return r.ResponseWriter.Write(b)
}

func checkSession(sess *quic.EarlySession) *quic.Session {
	qs, ok := (*sess).(quic.Session)
	if !ok {
		fmt.Println("Returned session is not quic sessions")
		return nil
	}
	return &qs
}

func checkFlowTeleSession(sess *quic.Session) *quic.FlowTeleSession {
	fs, ok := (*sess).(quic.FlowTeleSession)
	if !ok {
		fmt.Println("Returned session is not flowtele sessions")
		return nil
	}
	return &fs
}
