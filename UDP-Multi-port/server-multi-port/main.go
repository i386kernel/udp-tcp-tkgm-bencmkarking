package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func spinMultiPort(port string, quit chan struct{}) {
	udpAddr, err := net.ResolveUDPAddr("udp4", ":"+port)
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Started Concurrent UDP Server in IP %v Port %v", udpAddr.IP, udpAddr.Port)
	buffer := make([]byte, 1024)
	// _, err = new(net.UDPAddr), error(nil)
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
	var port = 10001
	quit := make(chan struct{})
	for i := port; i < (port + 10); i++ {
		go spinMultiPort(strconv.Itoa(i), quit)
	}
	<-quit // hang until an error
}
