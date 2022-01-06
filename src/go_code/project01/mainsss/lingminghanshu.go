package main

import "fmt"

//匿名函数


func main()  {
	//在定义匿名函数就直接调用，这种方式匿名函数只能调用一次
	re1 := func(n1 int,n2 int) int{
		r := n1+n2
		return r
	}(1,2)




	fmt.Println("匿名函数=",re1)


	//第二种用法
	re2 := func(n1 int,n2 int) int{
		r := n1-n2
		return r
	}

	fmt.Println("匿名函数re2=",re2(3,2))


	//全局匿名函数
	re3 :=Fun1(2,2)
	fmt.Println("全局匿名函数Fun1=",re3)



}

//全局匿名函数
var (
	Fun1 = func(n1 int,n2 int)( sum int){
		sum = n1*n2
		return
	}
)




