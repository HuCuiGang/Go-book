package singleton

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"unsafe"
)

type Singleon struct {
}

var singleInstance *Singleon
var once sync.Once

func GetSingletonObj() *Singleon {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleon)
	})
	return singleInstance
}

func GetSingle() *Singleon {
	if singleInstance != nil {
		return singleInstance
	}
	singleInstance = new(Singleon)
	return singleInstance
}

type on struct {
	o bool
	mu sync.Mutex
}

func  (n *on) Do(fn func())  {
	n.mu.Lock()
	defer func() {
		n.mu.Unlock()
	}()
	if n.o{
		return
	}
	fn()
	n.o=true
}

var oc on

func GetSingleSimple() *Singleon {
	oc.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleon)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x \n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestType1(t *testing.T)  {
	var i = "123"
	var i64 int64 = 34
	var i32 int32 = 5
	i64 =int64(i32)
	fmt.Println(i64)
	number, _ := strconv.ParseInt(i,10,64)
	fmt.Printf("%d %T",number,number)
}
