package main

import (
	"fmt"
	ioutil "io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*
1.发起百度搜索的get请求：“http://www.baidu.com/s?wd=肉丝”，打印出回复的内容
2.对https://httpbin.org/post 发起post请求，携带表单数据，键值自拟，打印回复内容
3.表单数据类型application/x-www-form-urlencoded , 表单读取API：string.NewReader("rmb=0.5&hobby=接客和约汉)
 */

func Get(){
	resp, err := http.Get("http://www.baidu.com/s?wd=肉丝")
	ClientHandError(err,"http.Get")

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	ClientHandError(err,"outil.ReadAll")

	fmt.Println(string(bytes))

}

func Post()  {

	resp, err := http.Post("https://httpbin.org/post?name=jack&teacher=john","application/x-www-form-urlencoded",strings.NewReader("rmb=0.5&hobby=接客和约汉"))
	ClientHandError(err,"ioutil.ReadAll")

	fmt.Println("123")

	bytes, err := ioutil.ReadAll(resp.Body)
	ClientHandError(err,"ioutil.ReadAll")

	fmt.Println(string(bytes))


}

func ClientHandError(err error,when string)  {
	if err!=nil {
		fmt.Println(err,when)
		os.Exit(1)
	}
}
