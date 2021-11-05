package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
/*
1.拨号连接服务器端主机的8888端口，建立UDP连接
2.循环从标准输入（命令行）读取一行用户输入，向服务端发送
3.接受并打印服务端消息，如果消息是：“bye”，就退出程序
*/

func ClientHandleError(err interface{}, when string) {
	if err != nil {
		fmt.Println(err,when)
		os.Exit(1)
	}
}

func main(){
	//解析得到的UDP地址
	conn,e :=net.Dial("udp","127.0.0.1:8888")
	ClientHandleError(e,"net.Dial(\"udp\",\"127.0.0.1\")")
	//准备标准输入（命令行）读取器
	reader :=bufio.NewReader(os.Stdin)
	//创建消息缓冲区
	buffer := make([]byte,1024)
	//循环发送消息
	for  {
		lineBytes,_,_ :=reader.ReadLine()
		conn.Write(lineBytes)
		n,_ := conn.Read(buffer)
		servMsg := string(buffer[:n])
		if servMsg =="fuckoff" {
			conn.Write([]byte("呜呼狡兔死飞鸟尽良弓藏，吾去也！"))
			break
		}
	}
	fmt.Println("Game Over!")
}
