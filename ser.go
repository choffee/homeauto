package main

import (
	"github.com/tarm/goserial"
	"log"
	"time"
	"io"
)


func serialReader(serialPort *io.ReadWriteCloser , line *chan<- *string ) {
	// Reads continuisly from a serial port and sends whole line back
	buf := make([]byte, 128)
	for {
		n, err := (*serialPort).Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Printf("%q", buf[n])
        }

}


func main() {
	c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	serialLines := new(chan<- *string)

	go serialReader(&s, serialLines)

	time.Sleep(2000 * time.Millisecond)

	_, err = s.Write([]byte("RF  A2off\n"))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2000 * time.Millisecond)

}
