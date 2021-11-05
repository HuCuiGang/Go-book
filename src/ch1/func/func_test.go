package _func

import (
	"fmt"
	"testing"
	"time"
)

func timeSpent(inner func(op int) int) func(op int )int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int{
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T) {
	 tsSf := timeSpent(slowFun )
	 t.Log(tsSf(10))
}

