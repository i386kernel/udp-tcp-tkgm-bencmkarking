package main

import (
	"fmt"
	"net"
)

const MaxBufferSize = 1024

//func WritetoFile(bs []byte) {
//	//newFile, err := os.Create("test.txt")
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//
//	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer file.Close()
//	file.Write(bs)
//
//}

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, MaxBufferSize)
	_, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("UDP client :", string(buffer))
	if err != nil {
		fmt.Println(err)
	}
	//write message back to client
	message := []byte("Hello UDP client!\n")
	_, err = conn.WriteToUDP(message, addr)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//hostname := "127.0.0.1"
	portNum := "10001"
	//service := hostname + ":" + portNum

	udpAddr, err := net.ResolveUDPAddr("udp4", ":"+portNum)
	if err != nil {
		fmt.Println(err)
	}
	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("UDP server up and listening on port :", portNum)
	defer ln.Close()
	for {
		handleUDPConnection(ln)
	}
}
