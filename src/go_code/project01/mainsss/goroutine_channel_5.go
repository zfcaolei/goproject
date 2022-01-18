package main

import "fmt"

//管道只读只写
func main()  {

	//管道申明为只读和只写
	//默认情况下管道是双向的

	//var chan1 chan int //双向管道

	//只写管道
	var chan2 chan <- int
	chan2 = make(chan int)
	fmt.Println(chan2)
	chan2<-200
	//b1 := <-chan2  //报错


	//只读管道
	var chan3 <-chan int
	chan3 = make(chan int)
	fmt.Println(chan3)
}
