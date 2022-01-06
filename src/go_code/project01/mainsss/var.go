package main

import "fmt"

//定义全局变量  可以定义不使用
var nan1 = "ddasdsa"
var nan2 = "ddasdsa"
var nan3 = "ddasdsa"

//多个全局变量定义
var (
	nan4 = "222"
	nan5 = "3333"
	nan6 = "66666"
)

func main() {

	//go变量使用方式
	//1,定义变量及类型不赋值有默认值
	var testname string
	testname = "cccccccccccccc"
	//2.根据值自行判断变量类型
	var name = "caolei"

	//3.省略var  注意左侧的变量不应该是声明过的，否则导致
	test := "caolllll"

	//4.多变量声明
	var name4, name5, name6 int
	var name1, name2, name3 int = 111, 222, 333

	var name7, name8 = 1111, "caolei"
	name9, name10 := 2222, "cccc"

	var (
		name11 = 11
		name22 = 22
	)

	//变量使用的注意事项
	//变量区域的值可以在同一类型下不断变化，但是不能改变类型
	var i int = 1
	i = 11
	i = 22
	//i = 2.2

	//var ss int
	ss := "cao"

	// 不能重复定义var i int = 3
	//i := 3

	fmt.Println("caolei", i, ss, name, testname, test, name1, name2, name3, name4, name5, name6, name7, name8, name9, name10, nan1, name11, name22)
}
