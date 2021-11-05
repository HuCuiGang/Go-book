package pipe_filter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToIntFilter struct {

}

func NewToIntFileter() *ToIntFilter {
	return &ToIntFilter{}
}

//将字符数组转换为int数组
func (tif *ToIntFilter) Process(data Request)(Response,error){
	parts,ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts{
		s,err := strconv.Atoi(part) //遍历传入的字符串数组 进行转换为int
		if err != nil  {
			return nil,err
		}
		ret = append(ret ,s) //将转换后的int 放入 int 数组
	}
	return ret,nil
}