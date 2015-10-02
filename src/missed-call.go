package main

import (
	"log"
	"github.com/tarm/serial"
	"time"
	"strings"
)

func main() {
	log.Printf("Hello Spammer!")


	c := &serial.Config{Name: "/dev/tty.HUAWEIMobile-Modem", Baud: 9600, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	defer s.Close()

	if err != nil {
		log.Fatal(err)
	}

	result := sendCommand(s,"ATE1", true)
	log.Printf("Reply: %q", result)

	//	buf := make([]byte, 2096)
	//	n, _ = s.Read(buf)
	//
	//	log.Printf("AT Reply: %q", buf[:n])
	//
	//	//n, err = s.Write([]byte("ATD7829869145;"))
	//	n, err = s.Write([]byte("AT"))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	n, _ = s.Read(buf)
	//	log.Printf("Dial Reply: %q", buf[:n])
	//
	//	log.Printf("Sleeping..")
	//	time.Sleep( time.Second * 5)
	//
	//	n, err = s.Write([]byte("AT+CHUP"))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	n, _ = s.Read(buf)
	//	log.Printf("Hang Reply: %q", buf[:n])


	log.Printf("Terminated")

}


func sendCommand(p *serial.Port, command string, waitForOk bool) string {
	log.Println("--- SendCommand: ", command)
	var status string = ""
	p.Flush()
	_, err := p.Write([]byte(command))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 32)
	var loop int = 1
	if waitForOk {
		loop = 10
	}
	for i := 0; i < loop; i++ {
		// ignoring error as EOF raises error on Linux
		n, _ := p.Read(buf)
		if n > 0 {
			status = string(buf[:n])
			log.Printf("SendCommand: rcvd %d bytes: %s\n", n, status)
			if strings.HasSuffix(status, "OK\r\n") || strings.HasSuffix(status, "ERROR\r\n") {
				break
			}
		}
	}
	return status
}