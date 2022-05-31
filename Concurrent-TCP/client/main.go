package main

import (
	"fmt"
	"net"
	"time"
)

//)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8081")
	//conn, err := net.DialTCP("tcp", nil, addr)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//start := time.Now().UnixMilli()
	//count := 0
	//for i := 0; i < 18000; i++ {
	//	fmt.Fprintf(conn, "from the client\n")
	//	message, _ := bufio.NewReader(conn).ReadString('\n')
	//	fmt.Print(message)
	//	count += 1
	//}
	//end := time.Now().UnixMilli()
	//diff := end - start
	//conn.Close()
	//fmt.Println(diff)
	//fmt.Println(count)
	start := time.Now().UnixMilli()
	for i := 0; i < 10; i++ {
		dialconn(addr)
	}
	end := time.Now().UnixMilli()
	diff := end - start
	fmt.Println(diff)
}

func dialconn(addr *net.TCPAddr) {
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println(err)
	}
	//	start := time.Now().UnixMilli()
	count := 0
	for i := 0; i < 13000; i++ {
		fmt.Fprintf(conn, "from the client\n")
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//	fmt.Print(message)
		count += 1
	}
	//end := time.Now().UnixMilli()
	//diff := end - start
	conn.Close()
	// fmt.Println(diff)
	fmt.Println(count)
}
