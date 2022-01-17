package main

import "fmt"

//channel管道关闭

func main()  {

	chan1 := make(chan int,3)

	chan1<- 100
	chan1<- 200

	close(chan1) //关闭管道 已经不能写入管道数据

	n1 := <-chan1	//关闭管道后但是能读出管道数据
	fmt.Println(n1)
}