package main

import (
	"fmt"
	"net"
)

//func listen(conn *net.UDPConn, quit chan struct{}) {
//	buffer := make([]byte, 1024)
//	n, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
//	for err == nil {
//		n, remoteAddr, err = conn.ReadFromUDP(buffer)
//		if err != nil {
//			fmt.Println(err)
//		}
//		log.Printf("From UDP Client %s || Message - %s\n", remoteAddr, string(buffer[:n]))
//	}
//	fmt.Println("listener failed - ", err)
//	quit <- struct{}{}
//}

//

func listen(conn *net.UDPConn, quit chan struct{}) {
	buffer := make([]byte, 1024)
	err := error(nil)
	for err == nil {
		_, _, err = conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("listener failed - ", err)
	quit <- struct{}{}
}

func main() {
	var port = "10001"
	// var ip = []byte{127, 0, 0, 1}

	udpAddr, err := net.ResolveUDPAddr("udp4", ":"+port)
	if err != nil {
		fmt.Println(err)
	}
	// addr := net.UDPAddr{Port: port, IP: ip}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	// log.Printf("Started Concurrent UDP Server in IP %v Port %v with %v Number of CPU's", udpAddr.IP, udpAddr.Port, numCPU)
	quit := make(chan struct{})
	for i := 0; i < 16; i++ {
		go listen(conn, quit)
	}
	<-quit // hang until an error
}
