package type_test

import (
	"errors"
	"testing"
)


func TestType(t *testing.T)  {
	x,y :=1 , 1
	t.Log(y)
	t.Log(x)
	for i:=0;i<10;i++ {
		temp := x
		x = y
		y = temp + x
		t.Log(y)
	}
}

func GetFibonacci(n int) ([]int,error) {
	if n<0 || n>100 {
		return nil,errors.New("n should be in [2,100]")
	}
	fibList := []int{1,1}
	for i := 2 ;i<n ; i++ {
		fibList = append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList,nil
}

func TestGetFibonacci(t *testing.T){
	if v,err:= GetFibonacci(-10);err!= nil {
		t.Error(err)
	}else {
		t.Log(v)
	}
}


func TestPoint(t *testing.T ){
	a:= 1
	artr := &a
	t.Log(a,artr)
	t.Logf("%T %T",a,artr)
}


