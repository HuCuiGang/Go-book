package pipe_filter

import (
	"fmt"
	"testing"
)

//测试割字符串Process 方法
func TestSplitFilter_Process(t *testing.T) {
	splitFilter := NewSplitFilter("-")
	part ,err := splitFilter.Process("1-2-3-4-5-6")
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(part)
}




