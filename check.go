package main

// module import
import (
	"fmt"
	"net"
	"time"
)

// Makes a call to the destination IP
func Check(destination string, port string) string {
	address := destination + ":" + port
	timeOut := time.Duration(5 + time.Second)
	conn, err := net.DialTimeout("tcp", address, timeOut)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable. \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is up and running.\n From: %v, To: %v", destination,
			conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
