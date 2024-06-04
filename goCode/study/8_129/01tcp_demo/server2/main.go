package main

import (
	"fmt"
	"net"
)

func server(conn net.Conn) {
	var tmp [1024]byte
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			return
		}
		go server(conn)
	}
}
