package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		println("dial error:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请发送信息：")
		tmp, _ := reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)
		if tmp == "exit" {
			break
		}
		conn.Write([]byte(tmp))
	}
	conn.Close()
}
