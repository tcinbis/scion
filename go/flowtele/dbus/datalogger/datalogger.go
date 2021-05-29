package datalogger

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/scionproto/scion/go/lib/log"
)

type channelData []string

type ChannelDataStruct interface {
	Strings() []string
}

type RTTData struct {
	FlowID    int
	Timestamp time.Time
	SRtt      time.Duration
}

type DbusDataLogger struct {
	writerChan chan channelData
	abortChan  chan struct{}
	csvFile    *os.File
	writer     *csv.Writer
	metaData   []string
	header     []string
}

func NewDbusDataLogger(csvFileName string, csvHeader []string, metaDataHeader []string) *DbusDataLogger {
	file, err := os.Create(csvFileName)
	check(err)
	d := DbusDataLogger{
		writerChan: make(chan channelData, 100),
		abortChan:  make(chan struct{}, 1),
		csvFile:    file,
		header:     append(csvHeader, metaDataHeader...),
	}
	d.writer = csv.NewWriter(d.csvFile)
	d.writeHeader()

	return &d
}

func (d *DbusDataLogger) SetMetadata(meta []string) {
	d.metaData = meta
}

func (d *DbusDataLogger) Send(data ChannelDataStruct) {
	d.writerChan <- data.Strings()
}

func (d *DbusDataLogger) SendString(data []string) {
	d.writerChan <- data
}

func (d *DbusDataLogger) writeHeader() {
	check(d.writer.Write(d.header))
	check(d.writer.Error())
}

func (d *DbusDataLogger) Close() {
	close(d.writerChan)
	close(d.abortChan)
}

func (d *DbusDataLogger) Run() {
	go func() {
		defer log.HandlePanic()
		defer func() {
			fmt.Println("### Closing csvFile ###")
			check(d.csvFile.Close())
			check(d.writer.Error())
		}()
	loop:
		for {
			select {
			case <-d.abortChan:
				fmt.Println("CSV Writer received abort. Flushing file")
				d.writer.Flush()
				check(d.writer.Error())
				fmt.Printf("%d left in writter channel\n", len(d.writerChan))
				break loop
			case s := <-d.writerChan:
				if len(s) > 0 {
					check(d.writer.Write(append(s, d.metaData...)))
					check(d.writer.Error())
				}
			}
		}
		fmt.Println("### Exiting datalogger run function ###")
	}()
}

func (r *RTTData) Strings() []string {
	return []string{strconv.Itoa(r.FlowID), strconv.Itoa(int(UnixMicroseconds(r.Timestamp))), strconv.Itoa(int(r.SRtt.Microseconds()))}
}

func UnixMicroseconds(t time.Time) int64 {
	return t.UnixNano() / 1e3
}

func check(err error) {
	if err != nil {
		fmt.Printf("Error in dbus datalogger: %v\n", err)
	}
}
