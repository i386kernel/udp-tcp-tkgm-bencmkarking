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

var numofclient = 8
var requestperclient = 125000
var addpo = "localhost:10001"
var numb = 32

func concurr(wg *sync.WaitGroup, addrstr *net.UDPAddr, numclient string, rqc int, pl string) {
	defer wg.Done()
	c, err := net.DialUDP("udp4", nil, addrstr)
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	for i := 0; i < rqc; i++ {
		msg := fmt.Sprintf("Client Number: %s, Message Number: %d, Message: %s", numclient, i, pl)
		_, err := c.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func createUDPClient() {

	numclient := flag.Int("nc", numofclient, "Define the number of clients")
	reqperclient := flag.Int("rc", requestperclient, "Define the messages per client")
	addrpo := flag.String("ap", addpo, "Define the address and Port -> 'address:port'")
	numbytes := flag.Int("nb", numb, "Define the number of bytes to be sent as payload, Payload Size")
	flag.Parse()

	fmt.Printf("Number of clients/goroutines establishing connection witn UDP Server %d\n\n", *numclient)
	fmt.Printf("Number of Requests per client %d\n\n", *reqperclient)
	fmt.Printf("Payload size with each request: %d\n", *numbytes)

	var buffer bytes.Buffer

	for i := 0; i < (*numbytes+1)-50; i++ {
		buffer.WriteString("a")
	}
	s, err := net.ResolveUDPAddr("udp4", *addrpo)
	if err != nil {
		fmt.Println(err)
	}
	start := time.Now().UnixMilli()
	for i := 0; i < *numclient; i++ {
		wg.Add(1)
		go concurr(&wg, s, strconv.Itoa(i), *reqperclient, buffer.String())
	}
	wg.Wait()
	end := time.Now().UnixMilli()
	diff := end - start
	microdiff := diff * int64(time.Microsecond)
	fmt.Println("Total time in microseconds would be: ", microdiff)
	totalrequests := *numclient * *reqperclient
	perrequest := float32(microdiff) / float32(totalrequests)

	fmt.Printf("Total Number of requests %d\n", totalrequests)
	fmt.Printf("Total Time took to send %d requests would be %d milliseconds\n", totalrequests, diff)
	fmt.Printf("It takes %f microseconds per request\n", perrequest)

}

func main() {
	createUDPClient()
}
