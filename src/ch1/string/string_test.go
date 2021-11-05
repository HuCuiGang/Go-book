package string

import (
	"strings"
	"testing"
)

func TestStringToRune(t *testing.T){
	s:="中华人民共和国"
	for _,c:=range s{
		t.Logf("%[1]c %[1]x",c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _,part := range parts{
		 t.Log(part)
	}
	t.Log(parts)
	t.Log(strings.Join(parts,"-"))
}

func Sum(ops ...int) int  {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T)  {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

