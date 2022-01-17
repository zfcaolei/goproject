package main

import "fmt"

//channel管道
func main()  {


	//定义channel管道  引用类型
	var testchanint chan int

	testchanint = make(chan int,3)  //容量要定义超出会报错

	fmt.Println(testchanint)

	//存入管道数据
	testchanint<-10

	num :=200
	testchanint<-num

	fmt.Println("testchanint管道长度=",len(testchanint),"testchanint容量=",cap(testchanint))


	//取出管道数据
	n1 := <-testchanint
	fmt.Println(n1)
	fmt.Println("testchanint管道长度=",len(testchanint),"testchanint容量=",cap(testchanint))


	//取出第二条数据 200
	n2 := <-testchanint
	fmt.Println(n2)


	//管道里面没有数据 再取出数据会报错  n3 := <-testchanint
	//取出数据不使用 <-testchanint



}