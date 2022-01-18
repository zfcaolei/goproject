package main

import (
	"fmt"
	"time"
)
//一个协程发生错误捕获错误不让其他协程产生影响退出
func sayhellosss()  {
	for i:=1;i<10;i++{
		time.Sleep(1*time.Second)
		fmt.Println("hello,word")
	}
}

func testss()  {
	defer func() {
		if err := recover(); err!=nil{
			fmt.Println("test()发生错误",err)
		}
	}()
	var mapMap map[int]string
	mapMap[0] = "golang" //error
}

func main()  {
	go sayhellosss()
	go testss()
	//
	for i:=1;i<=10;i++{
		fmt.Println("main()ok=",i)
		time.Sleep(time.Second)
	}
}
