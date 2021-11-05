package main

import (
	"fmt"
	"net"
	"os"
)

/*
1.服务器在本机的8888端口建立TCP监听
2.为接入的每一个客户端开辟一条独立的协成
3.循环接收客户端消息，不管客户说什么都自动回复“已阅xxx”
4.如果客户端说的是"im off ",则回复“bye” 并断开连接
*/

func main(){

	//服务器在本机的8888端口建立TCP监听
	listener,err :=net.Listen("tcp","127.0.0.1:8888")
	ServerHandleError(err,"net,listen")


	for  {
		//循环接入所有客户端，得到专线连接
		conn,e := listener.Accept()//阻塞的
		ServerHandleError(e,"listener.Accept()")

		//开一条独立的协成与该连接交互
		go ChatWith(conn)
	}


}

func ChatWith(conn net.Conn) {
	//创建消息的缓冲区
	buffer := make([]byte ,1024)

	//不断的收发消息
	for  {
		//一个完整的消息回合

		//读取客户端发送来的消息，存入缓冲区
		n,err:=conn.Read(buffer)  //n是长度，n个字节
		ServerHandleError(err,"conn.Read(buffer)")
		//转换为字符串输出
		clienMsg := string(buffer[0:n])
		fmt.Printf("收到%s的消息：%s \n",conn.RemoteAddr(),clienMsg)
		//回复客户端消息 1.已阅：  2.收到im off 后关闭
		if clienMsg!="im off"{

			conn.Write([]byte("已阅："+clienMsg))
		}else {
			//客户端
			conn.Write([]byte("bye！"))
			break
		}
	}
	conn.Close()
	fmt.Printf("客户端%s 已断开连接",conn.RemoteAddr())
}

func ServerHandleError(err error,when string) {
	if err != nil {
		fmt.Println(err,when)
		os.Exit(1)
	}
}
