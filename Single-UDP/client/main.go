package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
	//"strconv"
	//"time"
)

func main() {
	hostname := "localhost"
	portNum := "10001"

	addpo := hostname + ":" + portNum

	addrpo := flag.String("ap", addpo, "Define the address and Port -> 'address:port'")
	nummess := flag.Int("nm", 1, "Enter the Number of messages to be sent")
	flag.Parse()

	RemoteAddr, err := net.ResolveUDPAddr("udp", *addrpo)

	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	message := "Hello UDP server!"
	msg := []byte(message)
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println(err)
	}

	start := time.Now().UnixMilli()

	for i := 0; i < *nummess; i++ {
		msg := []byte(message + strconv.Itoa(i) + "\n")
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println(err)
		}
	}
	end := time.Now().UnixMilli()

	diff := end - start
	fmt.Println(diff)
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP Server :", addr)
	fmt.Println("Received from UDP server: ", string(buffer[:n]))

	fmt.Printf("Total Duration in Milli Seconds: %d\n", diff)

}
