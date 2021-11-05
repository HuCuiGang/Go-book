package array_test

import (
	"testing"
)

func TestArray(t *testing.T){
	arr := [...]int{1,2,3}
	for i:=0 ;i<len(arr) ;i++ {
		t.Log(arr[i])
	}
	for idx,e := range arr{
		t.Log(idx)
		t.Log(e)
	}
}

func TestSliceInit(t *testing.T){
	var s0 []int
	t.Log(len(s0),cap(s0))
	s0 = append(s0,1)
	t.Log(len(s0),cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1),cap(s1))

	s2 := make([]int,3,5)
	t.Log(len(s2),cap(s2))
	for _,v :=range s2{
		t.Log(v)
	}
}

func TestSliceGrowing(t *testing.T){
	s := []int{}
	for i:=0 ;i<10 ;i++ {
		s = append(s,i)
		t.Log(len(s),cap(s))
	}

}
func TestSliceShareMempry(t *testing.T){
	year := []string{"Jan","Feb","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}
	Q2 := year[3:6]
	t.Log(Q2,len(Q2),cap(Q2))
}

func TestInitMap(t *testing.T){
	m1:=map[int]int{1:1,2:4,3:9}
	t.Log(m1[2])
	t.Log("len m1=",len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Log("len m2",len(m2))
	m3 := make(map[int]int,10)
	t.Logf("len m3=%d",len(m3))

}
func TestAccessNotExistingKey(t *testing.T){
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2]=0
	t.Log(m1[2])
	if v,ex := m1[2] ;ex{
		t.Logf("Key 3`s value is %d",v)
	} else {
		t.Log("key 3 is not existing .")
	}
}
