package main

import "fmt"

//使用协程写入管道和协程取出管道
//使用两个channel管道  一个存入取出，一个判断主线程是否退出

func Write(wirte chan int)  {
	for i:=1;i<=50;i++ {
		wirte <- i
		fmt.Println("！！！！写入管道",i)
	}
	//使用完关闭管道
	close(wirte)
}

func Read(wirte chan int,exitchan chan bool)  {
	for{
		wirtedata,ok := <-wirte
		if !ok {
			exitchan <- true
			break
		}
		fmt.Println("读出管道=",wirtedata)
	}
	close(exitchan)
}


func main()  {
	//写入数据和读出数据管道
	ChanWrite := make(chan int,10)

	//主线程判断是否完成管道
	ExitChan := make(chan bool,1)

	go Write(ChanWrite)
	//go Read(ChanWrite,ExitChan)
	//循环判断主线程管道是否有数据否则再退出
	for  {
		_,ok := <-ExitChan
		if !ok {
			break
		}
	}
}
