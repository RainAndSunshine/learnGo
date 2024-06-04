package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 利用file自带的read()
func readFromFile1() {
	//*file(结构体),error
	fileObj, err := os.Open("./main.go")
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	//必须先写错误判断，因为当err不为nil的时候，就不会返回*File
	//如果先defer，就会导致nil.Close()，空指针是没有方法的，会panic

	//记得关闭文件
	defer fileObj.Close()

	//Read 读文件
	//写法1	直接定义切片
	//var temp = make([]byte, 128)
	//写法2	定义一个数组，后切片
	var temp [128]byte
	//文件不止128个字节，所以循环读
	for {
		n, err := fileObj.Read(temp[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed,err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节", n)
		fmt.Println(string(temp[:n]))
		//当文件都被读取完，n就读不到了
		if n < 128 {
			return
		}
	}
}

// 利用bufio这个包读文件
func readFromFileByBufio() {
	//*file(结构体),error
	fileObj, err := os.Open("./main.go")
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}

	//记得关闭文件
	defer fileObj.Close()

	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	//循环一行一行读取
	for {
		//一行一行读取，读到指定的'分隔符'，一般都是换行符，所以是一行一行读取
		line, err := reader.ReadString('\n') //注意这个分隔符是字符
		//如果到达文件尾
		if err == io.EOF {
			return
		}
		//如果出错
		if err != nil {
			fmt.Printf("read file failed,err:%v", err)
			return
		}
		//本身就有换行
		fmt.Print(line)
	}

}

// 使用ioutil包来读取数据
func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failded,err:%v\n", err)
	}
	//强制转换为字符串类型
	fmt.Println(string(ret))
}

func main() {
	//readFromFile1()
	//readFromFileByBufio()
	readFromFileByIoutil()
}
