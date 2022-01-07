package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int
}

func main() {

	//方式1
	var t0 Person
	t0.Name = "花花"
	t0.Age = 11
	fmt.Println(t0)

	//方式2
	t1 := Person{"曹磊", 20}
	fmt.Println(t1)

	//方式3
	var t2 *Person = new(Person)
	(*t2).Name = "嘟嘟"
	(*t2).Age = 12
	//简化也可以这样写
	t2.Name = "地方"
	t2.Age = 23
	fmt.Println(*t2)

	//方式4
	var t3 *Person = &Person{}
	t3.Name = "都的"
	t3.Age = 23
	fmt.Println(*t3)

	//装换json格式
	t3joson, _ := json.Marshal(t3)
	fmt.Println(string(t3joson))

	//使用结构体方法
	s3 := t1.test()
	fmt.Println(s3)

	s4 := t0.jisuan()
	fmt.Println(s4)

	//使用String方法自动调用
	str := Person{
		"dfsf", 11,
	}
	fmt.Println(&str)

}

//结构体方法的声明
func (p Person) test() string {
	return p.Name
}

func (p Person) jisuan() int {
	res := 0

	for i := 0; i <= 100; i++ {
		res += 100
	}
	return res
}

func (p Person) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", p.Name, p.Age)
	return str
}
