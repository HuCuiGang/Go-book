package pipe_filter

import (
	"fmt"
	"testing"
)

//测试traightPipeline的Process方法
func TestStraightPipeline_Process(t *testing.T) {
	spliter := NewSplitFilter("-")
	converter := NewToIntFileter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("p1",spliter,converter,sum)
	ret,err := sp.Process("1-2-6")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
	//if ret != 6 {
	//	t.Fatalf("The expected is 6 ,but the actual is %d",ret)
	//}
}
