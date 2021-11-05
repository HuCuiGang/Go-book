package pipe_filter

import (
	"fmt"
	"testing"
)
//测试求和的Process方法
func TestSumFilter_Process(t *testing.T) {
	sumfilter := NewSumFilter()
	ret,err := sumfilter.Process([]int{1,2,4,5,6})
	if err!= nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d %T",ret,ret)
}
