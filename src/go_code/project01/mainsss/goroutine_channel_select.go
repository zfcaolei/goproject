package main

import "fmt"

//select
func main()  {
	//传统的方法遍历管道时，如果不关闭会阻塞而导致deadlock
	//问题，在实际开发中，可能我们不好确定什么时候关闭该管道
	//可以使用select 方式可以解决

	charint := make(chan int,10)
	for i:=1;i<=10;i++ {
		charint<- i
	}

	chanstr := make(chan string,5)
	for i:=1;i<=5;i++ {
		str := "hello-word"+fmt.Sprintf("%d",i)
		chanstr<- str
	}

	for {
		select {
		case intdata := <-charint:
			fmt.Println(intdata)

		case strdata := <-chanstr:
			fmt.Println(strdata)
		default:
			fmt.Println("都取到数据，没有了")
			return
		}
	}

}
