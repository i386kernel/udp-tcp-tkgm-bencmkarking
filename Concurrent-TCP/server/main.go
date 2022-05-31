package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnnection(conn net.Conn) {
	count := 0
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Message Received:", message)
		count += 1
		//newMessage := strings.ToUpper(message)
		//conn.Write([]byte(newMessage + "\n"))
		fmt.Println(count)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Accept connection on port")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Calling handleConnection")
		go handleConnnection(conn)
	}
}

//
//func main() {
//	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:"+strconv.Itoa(ser))
//	for ser := 8000; ser < 8010; ser++ {
//		tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:"+strconv.Itoa(ser))
//		if err != nil {
//			fmt.Println(err)
//		}
//		wg.Add(1)
//		go listener(tcpAddr, &wg)
//	}
//	wg.Wait()
//}
//
//func listener(tcpAddr *net.TCPAddr, wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Println(tcpAddr.IP, tcpAddr.Port)
//	fmt.Println("Listening in Address: ", tcpAddr.String())
//	listener, err := net.ListenTCP("tcp", tcpAddr)
//	if err != nil {
//		fmt.Println(err)
//	}
//	for {
//		buff := make([]byte, 1024*2)
//		conn, err := listener.Accept()
//		if err != nil {
//			fmt.Println(err)
//		}
//		n, err := conn.Read(buff)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(string(buff))
//		fmt.Println("read", n)
//		//daytime := time.Now().String()
//		//conn.Write([]byte(daytime))
//		//conn.Close()
//	}
//}
