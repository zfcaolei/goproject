package main

import (
	"fmt"
	"time"
)

//开启协程写入管道2000条数据
//开启8个协程 取出管道2000条数据

func WriteMag(ChanWrite chan int)  {
	for i:=1;i<=80000;i++{
		ChanWrite <- i
	}
	//写入完毕关闭管道
	close(ChanWrite)
}

func QuchuMag(ChanWrite chan int,ChanDeposit chan map[int]int,ChanExit chan bool)  {
	for  {
		//取出写入管道数据
		data,ok := <-ChanWrite
		if !ok {
			break
		}
		//定义一个map传入map管道
		count := 0
		mapdata := make(map[int]int)
		for i:=1;i<=data;i++ {
			count +=data
		}
		mapdata[data] = count
		ChanDeposit <- mapdata
	}
		fmt.Println("取出完毕!!!!!!!!!!")
		ChanExit <- true
}



func main()  {

	//需要三条管道 1.写入数据管道 2.取出写入数据然后存入管道 3主线程判断协程是否完成管道

	ChanWrite := make(chan int,200)

	ChanDeposit := make(chan map[int]int,80000)

	ChanExit := make(chan bool,4)

	start := time.Now().Unix()
	//开启协程写入数据到管道
	go WriteMag(ChanWrite)

	//开启8个协程处理
	for i:=1;i<=8;i++ {
		go QuchuMag(ChanWrite,ChanDeposit,ChanExit)
	}

	go func() {
		for i:=1;i<=8;i++ {
			<-ChanExit
		}
		end := time.Now().Unix()
		fmt.Println("使用协程消耗时间=",end-start)
		close(ChanDeposit)
	}()

	for  {
		data,ok := <-ChanDeposit
		if !ok {
			break
		}
		fmt.Println(data)
	}

	fmt.Println("主进程完毕退出")







}
