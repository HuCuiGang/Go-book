# Go 的基本数据类型
- int string byte float bool

## 数据类型之间的转换
- float转int
float64（int64）                           //直接类型后面（）内填入需要转的数据类型进行转换

- string 转int 
value，error strconv.ParseInt(str,10,64)  //使用strconv 方式转换

- int 转string 
方式1：str = fmt.Sprintf("%d",int64(18))  //使用fmt.Sprintf（）方式进行转换
方式2：str = fmt.strconv.FormatInt(int64(18),10)    //使用strconv 方式转换 ， 10 表示10进制
方式3：str = strconv.Itoc(int(18))        //strconv.Ioc 将整数转换成 string类型

- float 转 string
str = fmt.Spring("%f",float(18.2324))
str = fmt.strconv.FormatFloat(float(2324),'f',10,64) // 'f' 格式 ，10 小数点保留10位，64 float64 数据类型

- bool 转 string
str = fmt.Spring("%t",bool(flase))
str = fmt.strconv.FormatBool(bool(false))

- byte 转 string
str = fmt.Spring("%c",byte('a'))

# 常用数据结构
- 数组 切片 map 

## 数组的使用

- 访问数组元素
数组名[下标] 

- 四种初始化数组的方式
方式1 ： var numArr1 [3]int = [3]int{1,2,3}
方式2 :  var numArr2 = [3]int{1,2,3}
方式3 ： var numArr3 = [...]int{1,2,3}   //[...] 是规定写法
方式4 ： var numArr4 = [...]int{1:800 , 0:900,2:999} //在指定的下标位置插入对应的数据

- 数组的遍历
方式1： 常规遍历
for i:=0 ; i<len(numArr1) ; i++ {
    fmt.Println(numArr1[i])
}

方式2： for - range 结构遍历
for index,value : = range numArr{
    fmt.Println("数组下标=",index,"值=",value)
}

## 切片的使用

- 用"make"声明切片
mslice := make([]int,5,10) //创建一个int类型的切片。元素值默认初始化为0 ，长度为5 ，容量为10
长度： 切片当前有多少个元素
容量： 切片可以容纳多少个元素

- 用[]声明切片
slice1 := []int{1,2,3,4}  //[]type用于声明切片，[…]type用于声明数组。

- 引用元素
mslice[2] = 4 
mslice = append(mslice,8) //在mslice的末尾新增元素8.该操作将导致len+1；如果append之前cap就已经等于len的话，那么append还会导致mslice的cap按两倍扩展.
mslice[:]		//[:] 表明mslice的全部元素，输出：[0 0 4 0 0 8]
mslice[2:]		//[2:] 表明mslice下标为2开始的元素，输出：[4 0 0 8]
mslice[2:4]		//[2:4] 表明mslice下标为2~3的元素，输出：[4 0]
mslice[:4]		//[:4] 表明mslice下标为0~3的元素，输出：[0 0 4 0]

- 在旧切片的基础上生成新的切片
newSlice := mslice[2:4:6] //在mslice的基础上生成新的切片 -- newSlice的元素从mslice[2]开始。newSlice的len为4-2=2，cap为6-2=4。mslice[i:j:k]，其计算公式为len=j-i,cap=k-i。

## map的使用
- 申明和初始化
方法1：
map[KeyType]ValueType
var m map[string]int
m = make(map[string]int)

方法2：
var m map[string]int = map[string]int{"hunter":10 ,"tony":12}
m = map[string]int{}

- 读取
i:=m["route"] // 如果route存在，就返回那个值，如果不存在，返回0值，也就是说，根据这个value的类型，返回缺省值，比如string，就返回“”，int 就返回0

i, ok := m["route"] 

- 删除
delete[m,"route"] //如果route存在，删除成功，否则什么都没有发生

- 遍历
for key , value := range m {
    fmt.Println("key:"key , "value":value)
}
