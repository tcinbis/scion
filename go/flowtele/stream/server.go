package main

import (
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/flowtele"
	"github.com/netsec-ethz/scion-apps/pkg/shttp"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
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
	ip       = kingpin.Flag("ip", "ip to listen on").Default("0.0.0.0").String()
	port     = kingpin.Flag("port", "port the server listens on").Default("8001").Uint()
	useScion = kingpin.Flag("scion", "Enable serving server via SCION").Default("false").Bool()
)

func init() {
	kingpin.Parse()
	fmt.Printf("Parsing complete. Listening on %v:%v\n", *ip, *port)
}

func startTCPServer(handler http.Handler) {
	fmt.Printf("Using TCP\n")
	http.Handle("/", handler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *ip, *port), nil)
}

func startSCIONServer(handler http.Handler) {
	fmt.Printf("Using SCION\n")

	newSrttMeasurement := func(t time.Time, srtt time.Duration) {
		//fmt.Printf("t: %v, srtt: %v\n", t, srtt)
	}

	flowteleSignalInterface := flowtele.CreateFlowteleSignalInterface(newSrttMeasurement, func(t time.Time, newSlowStartThreshold uint64) {

	}, func(t time.Time, congestionWindow uint64, packetsInFlight uint64, ackedBytes uint64) {

	})
	// make QUIC idle timout long to allow a delay between starting the listeners and the senders
	quicConfig := &quic.Config{
		//MaxIdleTimeout: time.Second,
		//KeepAlive: true,
		FlowTeleSignal: flowteleSignalInterface,
	}

	server := shttp.NewScionServer(fmt.Sprintf("%s:%d", *ip, *port), withLogger(handler), nil, quicConfig)
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
