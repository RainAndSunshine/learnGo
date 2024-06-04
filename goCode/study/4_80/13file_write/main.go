package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//文件写入操作
//os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能
//func OpenFile(name string,flag int,perm FileMode)(*File,err){
//...
//}
//其中：name(要打开的文件名)、flag(打开文件的模式)、perm(文件权限，一个八进制数，Linux常用,Windows没用)

//文件模式
//只写		os.O_WRONLY
//创建文件	os.O_CREATE
//只读		os.O_RDONLY
//读写		os.O_RDWR
//清空		os.O_TRUNC
//追加		os.O_APPEND

//perm(文件权限，一个八进制数)
//r(读)04
//w(写)02
//x(执行)01

//打开文件写内容

func writeDemo1() {
	//'|'相当于二进制的或运算
	fileObj, err := os.OpenFile("./xx.html", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()

	//Write
	fileObj.Write([]byte("周林懵逼了！\n"))
	//WriteString
	fileObj.WriteString("周林解释不了！\n")
}

// 利用bufio对文件进行操作
// 会将数据先保存至缓冲区中，最后通过Flush()刷盘，上传到磁盘或者代码中
func writeDemo2() {
	//'|'相当于二进制的或运算
	fileObj, err := os.OpenFile("./xx.html", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("Hello World\n") //将数据先写入缓存
	wr.Flush()                      //将缓存中的内容写入文件
}

func writeDemo3() {
	str := "陕西科技大学司马"
	err := ioutil.WriteFile("./xx.html", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}

func main() {
	//writeDemo1()
	//writeDemo2()
	writeDemo3()
}
