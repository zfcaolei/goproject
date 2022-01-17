package main

import "fmt"

func main()  {
	//定义一个管道，存入三条数据，然后再取出

	chantest := make(chan int,3)

	n1 := 100
	chantest <-n1
	chantest <-200
	chantest <-300
	fmt.Println("管道长度=",len(chantest))

	//取出

	t1 := <-chantest
	t2 := <-chantest
	t3 := <-chantest

	fmt.Println(t1,t2,t3)
	fmt.Println("管道长度=",len(chantest))


	//map数据类型插入管道
	chantest2 := make(chan map[string]string,4)

	v1 := make(map[string]string)
	v1["name"] = "曹磊"
	v1["age"] = "18"
	chantest2 <- v1

	v2 := map[string]string{"name":"王鑫","age":"33"}
	chantest2 <- v2

	chantest2 <- map[string]string{"name":"xxx","age":"332"}

	v11 := <-chantest2
	fmt.Println(v11)

}
