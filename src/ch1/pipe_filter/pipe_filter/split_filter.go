package pipe_filter

import (
	"errors"
	"strings"
)

//错误类型
var SplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct{
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

//切割字符串
func (sf *SplitFilter) Process(data Request) (Response,error)  {
	str ,ok := data.(string) //判断数据类型是否是可处理的类型
	if !ok {
		return nil , SplitFilterWrongFormatError
	}
	parts := strings.Split(str,sf.delimiter)
	return parts,nil
}