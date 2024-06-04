package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

// 参数不能少
func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("xx.html")
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	//对于Get请求，参数都放在URL上(query param)，请求体中是没有数据的
	queryParam := r.URL.Query()    //自动帮我们识别URL中的query param
	name := queryParam.Get("name") //调用里面的方法，根据key取value
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) //我在服务端打印客户端发来的请求Body
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
