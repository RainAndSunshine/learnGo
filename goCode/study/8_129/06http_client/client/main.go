package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client

func main() {
	//内置的请求
	//前面是URL   后面用问号分隔的是参数，用等于号分隔key和value
	//resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=科大&age=20")
	//if err != nil {
	//	fmt.Println("get url failed,err", err)
	//	return
	//}

	//自己造一个请求
	data := url.Values{} //url encode
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "周林?=") //遇到这种情况，必须encode，不然会冲突
	data.Set("age", "9000")
	urlStr := data.Encode() //URL encode之后的URL
	fmt.Println(urlStr)
	urlObj.RawQuery = urlStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("get url failed,err", err)
		return
	}
	//从resp中把服务端返回的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body) //我在客户端读出服务端返回的Body
	if err != nil {
		fmt.Println("read resp.Body failed,err", err)
		return
	}
	fmt.Println(string(b))
}
