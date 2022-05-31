package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
	}
	start := time.Now().UnixMilli()

	message := []byte("Hi TCP Server")
	for i := 0; i < 50000; i++ {
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println(err)
		}
	}
	end := time.Now().UnixMilli()
	diff := end - start
	fmt.Println(diff)
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
