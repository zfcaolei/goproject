package main

import "fmt"

type cat struct {
	Name string
	Age int
}

func main()  {

	chan1 := make(chan interface{},3)

	chan1<- "test1"

	chan1<- 22222

	chan1<- cat{Name: "ccc",Age: 23}



	//直接从管道取出最后一个结构体类型

	<-chan1
	<-chan1

	cat1 := <-chan1
	fmt.Printf("cat1类型是=%T cat1值是=%v",cat1,cat1)

	//但是这样取出会报错
	//fmt.Printf(cat1.Name)
	//因为chan1 是管道存入是接口类型了，这样取出要进行类型断言
	cat11 := cat1.(cat)
	fmt.Println(cat11.Name)
}
