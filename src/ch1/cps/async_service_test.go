package cps

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond *500)
	return "Done"
}

func otherTask()  {
	fmt.Println("working on something else")
	fmt.Println("Task is done")
}

func TestService(t *testing.T){
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string,1)
	//retCh := make(chan string,1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		//time.Sleep(time.Second * 10)
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsynSerVice(t *testing.T)  {
	retCH := AsyncService()
	otherTask()
	fmt.Println(<-retCH)
}

//多路选择 ，超时机制
//先接受到哪个channl的消息就先执行哪个channl后面的代码
func TestSelet(t *testing.T)  {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Second): //超时机制
		t.Error("time out")
	}
}
