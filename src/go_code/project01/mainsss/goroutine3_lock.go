package main

import (
	"fmt"
	"sync"
	"time"
)

//模拟并发200个协程写入map中错误

var (
	maps = make(map[int]int,10)
	lock sync.Mutex
)

//lock是一个全局的互斥锁


func test001(n int){
	count := 0
	for i:=0;i<=n;i++ {
		count += i
	}
	lock.Lock() //加锁
	maps[n]=count
	lock.Unlock() //解锁

}

func main()  {

	for i:=1;i<=200;i++ {
		go test001(i)
	}

	time.Sleep(time.Second * 10)
	lock.Lock() //加锁
	for k,v := range maps{
		fmt.Println("k=",k,"v=",v)
	}
	lock.Unlock() //解锁


}