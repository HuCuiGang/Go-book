package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
拨号连接服务端主机8888端口，建立连接
循环从标准输入（命令行）读取一行用户输入，想服务端发送
接受并打印服务端消息，若果消息是“bye”，就退出程序
*/


func main() {
	//拨号连接服务端主机8888端口，建立连接
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	ClientHandleError(err,"net.Dial()")

	//创建消息缓冲区
	buffer := make([]byte,1024)
	//准备标准输入（命令行）读取器
	reader := bufio.NewReader(os.Stdin)
	//不断的收发消息
	for{
		//接收命令行输入的 一行消息
		lineBytes , _, _:=reader.ReadLine()
	    //向服务端发送消息
		conn.Write(lineBytes)
		//接收服务端消息，长度为N个字节，
		n,err := conn.Read(buffer) //n 字节长度
		ClientHandleError(err,"conn.Read(buffer)")
		//转换为字符串
		serverMsg:=string(buffer[0:n])
		fmt.Println("服务端：",serverMsg)
		//若果收到服务端说bye就退出消息循环
		if serverMsg =="bye"{
			break
		}
	}
	//客户端程序结束
	fmt.Println("GAME OVER!")
}

func ClientHandleError(err interface{}, when string) {
	if err != nil {
		fmt.Println(err,when)
		os.Exit(1)
	}
}
