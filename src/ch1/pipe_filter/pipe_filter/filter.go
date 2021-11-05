package pipe_filter
//接收参数
type Request interface {}
//返回结果
type Response interface{}

//Filter类型
type Filter interface{
	Process(data Request)(Response,error)
}
