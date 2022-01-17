package main

import "fmt"

//channel管道变量

func main()  {

	chan1 := make(chan int,100)
	//向管道插入100条数据
	for i:=1;i<=100;i++ {
		chan1 <- i*2
	}

	//遍历管道取出一定使用for rangge 来遍历，而且遍历前一定要关闭管道否则会报错
	close(chan1)

	for v :=range chan1{
		n1 := v
		fmt.Println(n1)
	}

	fmt.Println(len(chan1))
	
}