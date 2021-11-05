package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc(
		//路由
		"/whoami",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("请求方法="+request.Method))
			writer.Write([]byte("请求c长度="+strconv.Itoa(int(request.ContentLength))))
			writer.Write([]byte("请求方法="+request.Host))
			writer.Write([]byte("请求方法="+request.Proto))
			writer.Write([]byte("请求方法="+request.RemoteAddr))
			writer.Write([]byte("请求的路由="+request.RequestURI))
			writer.Write([]byte("朕已收到请求"))

		},
	)

	http.HandleFunc("/baidu", func(writer http.ResponseWriter, request *http.Request) {
		baiduBytes, err := ioutil.ReadFile("/home/cg/go/src/ch1/HTTP/百度一下，你就知道.html")
		ClientHandError(err,"ioutil.ReadFile")
		writer.Write(baiduBytes)
	})

	http.ListenAndServe("127.0.0.1:8080",nil)
}