package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	service := ":8000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
