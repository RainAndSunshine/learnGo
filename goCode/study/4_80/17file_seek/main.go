package main

import (
	"fmt"
	"io"
	"os"
)

// 在文件中间读取数据和写入数据

// 读取
func f1() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	//光标移动
	//(偏移量    相对位置：0-文件头  1-当前位置  2-文件尾)
	fileObj.Seek(3, 0) //第一个字母a，然后window换行符是\r\n两个，所以偏移量为3

	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed,err:%v\n", err)
		return
	}
	fmt.Println(string(ret[:n]))
}

// 直接写入(不过会覆盖)
func f2() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()

	fileObj.Seek(3, 0)
	//这样写的话会覆盖掉原来位置的值，如果需要插入的话，需要新建一个临时文件
	fileObj.Write([]byte{'s'})
}

// 插入数据(新建临时文件)
// 因为没有办法直接在文件中间插入内容，所以需要借助一个临时文件
func f3() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()

	//文件创建
	//tmpObj, err := os.Create("./tep.txt")

	//直接打开的时候创建也可以
	tmpObj, err := os.OpenFile("./tmp.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("create file failed,err:%v\n", err)
		return
	}
	defer tmpObj.Close()
	//读取文件，写入临时文件
	var s [1]byte
	n, err := fileObj.Read(s[:])
	if err != nil {
		fmt.Printf("read from file failed,err:%v\n", err)
	}
	//写入临时文件
	tmpObj.Write(s[:n])
	//再写入要插入的内容
	var s1 = []byte{'c'}
	tmpObj.Write(s1)
	//紧接着把原文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		tmpObj.Write(x[:n])
	}
	//原文件后续的也写入了临时文件中，关闭两个文件
	fileObj.Close()
	tmpObj.Close()
	//将临时文件改名为原文件，替代原文件
	os.Rename("./tmp.txt", "./sb.txt")
}

func main() {
	//f1()
	//f2()
	f3()
}
