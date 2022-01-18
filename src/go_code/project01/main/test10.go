package main

import (
	"fmt"
	"time"
)

//开启协程写入管道2000条数据
//开启8个协程 取出管道2000条数据

func WriteMag(ChanWrite chan int)  {
	for i:=1;i<=8000;i++{
		ChanWrite <- i
	}
	fmt.Println("写入成功")

}

func QuchuMag(ChanWrite chan int,ChanDeposit chan map[int]int)  {
	for  {
		//取出写入管道数据
		select {
			case data,_:= <-ChanWrite:
				//定义一个map传入map管道
				count := 0
				mapdata := make(map[int]int)
				for i:=1;i<=data;i++ {
					count +=data
				}
				mapdata[data] = count
				ChanDeposit <- mapdata
			default:
				fmt.Println("取出完毕!!!!!!!!!!")
				return
		}
	}
}
func Suck(ChanDeposit chan map[int]int)  {
	for  {
		select {
			case v := <-ChanDeposit:
				fmt.Println(v)
			default:
				fmt.Println("无数据！！取完")
				return
		}
	}
}



func main()  {

	//需要三条管道 1.写入数据管道 2.取出写入数据然后存入管道 3主线程判断协程是否完成管道

	ChanWrite := make(chan int,100)

	ChanDeposit := make(chan map[int]int,8000)


	//开启协程写入数据到管道
	go WriteMag(ChanWrite)


	//开启8个协程处理
	for i:=1;i<=8;i++ {
	   go QuchuMag(ChanWrite,ChanDeposit)
	}
	time.Sleep(30*time.Second)

	Suck(ChanDeposit)

	//for i:=1;i<=800;i++ {
	//	fmt.Println(i)
	//}

}
