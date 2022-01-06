package main

import (
	"fmt"
	"strings"
)

//闭包函数
func res()func(int) (int,string) {
	n := 0
	str := "hello"
	return func(n1 int) ( int,string)  {
		n = n1+n
		str += "a"
		return  n,str
	}
}

func bibao() func(n1 int)int {
	n :=0
	return func(n1 int)int {
		n += n1
		return  n
	}
}

//strings.HasSuffix
//体验一下闭包的好处，如果用传统的函数也可以轻松的实现，但是需要每次都传入后缀名称，
//比如jpg,而闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复使用
func makeSuffix(suffix string) func( string) string {

	return func(name string) string {
		if !strings.HasSuffix(name,suffix){
			return  name+"."+suffix
		}
		return  name
	}
}

//不用闭包写法
func makeSuffix2(name string,suffix string,)string  {
	if !strings.HasSuffix(name,suffix){
		return  name+"."+suffix
	}
	return name
}





func main(){


	f1 := res()

	fmt.Println(f1(1))
	fmt.Println(f1(2))
	fmt.Println(f1(3))


	f2 := makeSuffix("jpg") //如果使用闭包完成，好处是后缀只要传一次就是
	fmt.Println("使用闭包函数判断后缀名makeSuffix=",f2("4fdsf.png"))
	fmt.Println("使用闭包函数判断后缀名makeSuffix=",f2("4fdsf.jpg"))


	fmt.Println(makeSuffix2("window.png","jpg"))
}