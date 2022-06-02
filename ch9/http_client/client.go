package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/hello/?name=xfp&age=23")
	// if err != nil {
	// 	fmt.Println("get url failed ,err :", err)
	// 	return
	// }
	//短连接即将DisableKeepAlives置于true，关闭keep alive
	// tr := &http.Transport{
	// 	DisableKeepAlives: true,
	// }
	// client := http.Client{
	// 	Transport: tr,
	// }
	//长连接适用于请求不太频繁的，用完就关闭
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/hello/")
	data.Set("name", "xfp")
	data.Set("age", "22")
	urlQuery := data.Encode()
	urlObj.RawQuery = urlQuery
	rep, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(rep)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//发请求
	//从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body) //在客户端读出服务端响应的body
	if err != nil {
		fmt.Println("read from resp failed ,err :", err)
		return
	}
	fmt.Println(string(b))
}
