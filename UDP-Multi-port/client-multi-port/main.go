package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

var numofclient = 10
var requestperclient = 100
var addpo = "localhost"
var numb = 128
var startport = 10001

func connectmulticlient(wg *sync.WaitGroup, udpaddr, msg string, rqc, clientnum int) {
	defer wg.Done()
	addrstr, err := net.ResolveUDPAddr("udp4", udpaddr)
	if err != nil {
		fmt.Println(err)
	}
	c, err := net.DialUDP("udp4", nil, addrstr)
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	for i := 0; i < rqc; i++ {
		msg := fmt.Sprintf("Client Number: %v, Message Number: %d, Message: %s", clientnum, i, msg)
		_, err := c.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func createUDPClient() {
	//	numclient := flag.Int("nc", numofclient, "Define the number of clients")
	reqperclient := flag.Int("rc", requestperclient, "Define the messages per client")
	addrpo := flag.String("ap", addpo, "Define the address and Port -> 'address:port'")
	numbytes := flag.Int("nb", numb, "Define the number of bytes to be sent as payload, Payload Size")
	flag.Parse()

	fmt.Printf("Number of Requests per client %d\n\n", *reqperclient)
	fmt.Printf("Payload size with each request: %d\n", *numbytes)

	var buffer bytes.Buffer
	for i := 0; i < (*numbytes+1)-50; i++ {
		buffer.WriteString("a")
	}
	fmt.Println("Payload Data would be: ", buffer.String())
	start := time.Now().UnixMilli()
	for i := 10001; i < 10011; i++ {
		wg.Add(1)
		go connectmulticlient(&wg, *addrpo+":"+(strconv.Itoa(i)), buffer.String(), *reqperclient, i)
	}
	wg.Wait()
	end := time.Now().UnixMilli()
	diff := end - start
	microdiff := diff * int64(time.Microsecond)
	fmt.Println("Total time in microseconds would be: ", microdiff)
	totalrequests := *reqperclient * 10
	perrequest := float32(microdiff) / float32(totalrequests)

	fmt.Printf("Total Number of requests %d\n", totalrequests)
	fmt.Printf("Total Time took to send %d requests would be %d milliseconds\n", totalrequests, diff)
	fmt.Printf("It takes %f microseconds per request\n", perrequest)
}

func main() {
	createUDPClient()
}
