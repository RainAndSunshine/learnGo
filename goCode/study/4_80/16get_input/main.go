package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入时，如果有空格

// 用Scan会出现错误，遇到空格就停止输入了
func useScan() {
	var s string
	fmt.Print("请输入内容:")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是:%s\n", s)
}

func useBufio() {
	var s string
	//在系统的输入里面读取，属于标准输入这个文件
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容:%s\n", s)
}

func main() {
	//useScan()
	useBufio()
}
