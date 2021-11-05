package pipe_filter

import (
	"fmt"
	"testing"
)
//测试转换为int数组的Process方法。
func TestToIntFilter_Process(t *testing.T) {
	tointfilter := NewToIntFileter()
	str := []string{"1","2","3"}
	ret,err := tointfilter.Process(str)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d %T",ret,ret)
}
