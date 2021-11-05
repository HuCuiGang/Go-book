package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
	"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T)  {
	e := new (Employee)
	err := json.Unmarshal([]byte(jsonStr),e) //将json字符串传入对象里
	if err != nil{
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e);err == nil{ //将对象按json格式解析
		fmt.Println(string(v))
	}else{
		t.Error(err)
	}

}
