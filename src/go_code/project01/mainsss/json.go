package main

import (
	"encoding/json"
	"fmt"
)
// json格式序列号
type Peopledata struct {
	Name string  `josn:"name"`
	Age int

}
//结构体序列号json格式
func teststruct()  {
	people1 :=Peopledata{
		Name: "曹磊",
		Age: 22,
	}

	data,err := json.Marshal(people1)
	if err !=nil{
		fmt.Sprintln("json序列化失败 err=",err)
	}

	fmt.Printf(string(data))
}

//map 格式化json数据
func testmap()  {

	v1 := map[string]interface{}{}
	v1["name"] = "曹磊"
	v1["age"] = 21
	data,err := json.Marshal(v1)
	if err !=nil{
		fmt.Sprintln("json序列化失败 err=",err)
	}

	fmt.Printf(string(data))
}

//slice 切片 格式化json数据类型
func testslice()  {
	var v2 []map[string]interface{}

	m1 := map[string]interface{}{"name":"堡垒","age":2}
	v2 = append(v2,m1)

	m2 := map[string]interface{}{"name":"堡垒1","age":21}
	v2 = append(v2,m2)

	data,err := json.Marshal(v2)
	if err !=nil{
		fmt.Sprintln("json序列化失败 err=",err)
	}

	fmt.Printf(string(data))
}
//对基本类型也可以序列号 但是意义不大
func testfloat64()  {
	num1 := 23.23
	data,err := json.Marshal(num1)
	if err !=nil{
		fmt.Sprintln("json序列化失败 err=",err)
	}
	fmt.Printf(string(data))
}

func main()  {
	teststruct()
	//testmap()
	//testslice()
	//testfloat64()
}