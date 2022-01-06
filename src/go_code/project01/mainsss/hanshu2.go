package main

import (
	"fmt"
)

func maxsun(n1 int,n2 int)int   {
	return  n1+n2
}




//函数也是一种类型可以赋值给变量，然而该变量也是一种函数，可以使用该变量调用该函数
func main()  {

	a := maxsun

	result :=a(1,20)

	fmt.Println(result)


	fmt.Println("--------------------")

	res := functions(a,1,33)
	fmt.Println(res)

	res2 := functions2(a,32,34)
	fmt.Println(res2)

	fmt.Println("--------------------")

	//可以给int取别名，再go中ss 和int 虽然都是int类型，但是go认为ss 和 int是两个类型
	type ss int

	var num5 ss
	var num6 int = 55
	num5 = ss(num6)

	fmt.Println(num5)


	type MyInt1 int
	type MyInt2 = int


	var i int =0
	var i1 MyInt1 = MyInt1(i) //error
	var i2 MyInt2 = i
	fmt.Println("i1=",i1,"i2=",i2)


	fmt.Println("______________________________")

	a3,a5 :=test7(23,43)
	fmt.Println("a3=",a3,"a5=",a5)


	//使用_下划线省略返回值
	a4,_ :=test7(23,43)
	fmt.Println("a4=",a4)

	fmt.Println("______________________________")


	test10 := test9(12,23)
	fmt.Println("可变参数函数test9=",test10)

}

type funsage func(int,int) int

//函数也可以作为形参，并且调用
func functions(abc func(int,int) int,num1 int,num2 int) int {
	return abc(num1,num2)
}
//函数形参也可以设置别名
func functions2(abc funsage,num1 int,num2 int) int {
	return abc(num1,num2)
}


//支持函数对返回值命名

func test7 (b1 int,b2 int)(sum int,sub int)  {
	sum = b1+b2
	sub = b1-b2
	return
}


//函数可变参数的使用
//可变形参要放在最后，函数也可以有一个可变形参
func test9(n1 int ,moreage ... int)  int  {
	sum  :=n1
	//遍历moreage参数
	//for i:=0;i<len(moreage);i++ {
	//	sum += moreage[i]
	//}

	for _,val :=range moreage{   //上下两种循环都可以
		sum += val
	}
	return sum
}