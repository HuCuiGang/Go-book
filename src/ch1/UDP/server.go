package main

import (
	"fmt"
	"net"
	"os"
)

/*
 1.服务端在本机的8888端口建立UDP监听
 2.循环接收客户端消息，不管客户端说什么，都自动回复“已阅xxx”
 3.如果客户端说的是“im off”，则回复“bye”
 */

func ServertHandleError(err interface{}, when string) {
	if err != nil {
		fmt.Println(err,when)
		os.Exit(1)
	}
}
func main() {
	//解析得到的UDP地址
	udpAddr,err :=net.ResolveUDPAddr("udp","127.0.0.1:8888")
	ServertHandleError(err,"et.ResolveUDPAddr(\"udp\",\"127.0.0.1:8888\")")
	//建立UDP监听，得到广口链接
	udpConn,err := net.ListenUDP("udp",udpAddr)
	ServertHandleError(err,"net.ListenUDP(\"udp\",udpAddr)")
	//创建消息缓冲区
	buffer := make([]byte,1024)
	//从广口连接中不断的读取（来自任何客户端）的数据包
	for  {
		//读取一个数据包到消息缓冲区，同时获得该数据包的客户端信息
		n,remoteAddr,err := udpConn.ReadFromUDP(buffer)
		ServertHandleError(err,"dpConn.ReadFromUDP(buffer)")
		//打印数据包消息内容
		clientMsg := string(buffer[:n])
		fmt.Printf("收到来自%v的消息：%s",remoteAddr,clientMsg)
		//回复该数据包的客户端
		if clientMsg !="im off"{
			_, err = udpConn.WriteToUDP([]byte("已阅："+clientMsg), remoteAddr)
			ServertHandleError(err,"WriteToUDP")
		}else {
			udpConn.WriteToUDP([]byte("fuckoff"), remoteAddr)
		}
	}
}
