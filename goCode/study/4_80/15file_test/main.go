package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//实现文件的复制
//1.先在文件中读取
//2.创建新文件，将数据写入文件

func fileCopy1() {
	fileObj, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()

	fileTemp, err := os.OpenFile("./write.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileTemp.Close()

	//创建一个reader对象
	reader := bufio.NewReader(fileObj)
	//创建一个writer对象
	writer := bufio.NewWriter(fileTemp)
	for {

		//一行一行读，以'\n'为分割符
		line, err := reader.ReadString('\n')
		//将数据写入缓存中
		writer.WriteString(line)
		//将缓存中的数据写进文件
		writer.Flush()
		//如果读到文件末尾
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed,err:%v\n", err)
			return
		}
	}
}

// 使用标准库里的io.Copy()函数
func fileCopy2(dstName, srcName string) (written int64, err error) {
	//以只读的方式打开文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer dst.Close()
	//io.Copy()
	return io.Copy(dst, src) //调用io.Copy()函数拷贝内容
}

func main() {
	//fileCopy1()
	_, err := fileCopy2("dst.txt", "src.txt")
	if err != nil {
		fmt.Printf("file copy failed,err:%v\n", err)
		return
	}
	fmt.Println("copy file success")
}
