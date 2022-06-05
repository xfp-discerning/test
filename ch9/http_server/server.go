package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//net/http server
func f1(w http.ResponseWriter, r *http.Request) {
	str := "hello world!"
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) { //在服务端打印客户端发来的请求的body
	// fmt.Println(r.URL)
	//对于get请求，参数都放在URL上（query，Param），请求体中是没有数据的
	queryParam := r.URL.Query() //自动帮我们识别URL中参数
	name := queryParam.Get("name")
	param := queryParam.Get("age")
	fmt.Println(name)
	fmt.Println(param)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/h/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
