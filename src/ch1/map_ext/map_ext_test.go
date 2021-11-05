package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T){
	m := map[int]func(op int) int {}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op*op}
	m[3] = func(op int) int { return op*op*op  }
	for key,v :=range m{
		t.Logf("key = %d,value= %d",key,v(2))
	}
}

func TestMapForSet(t *testing.T){
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is not exiting",n)
	}
	mySet[3] = true
	mySet[4] = false
	t.Log("mySet len = ",len(mySet))
	delete(mySet,1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is not exiting",n)
	}
}
