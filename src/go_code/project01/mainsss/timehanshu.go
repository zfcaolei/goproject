package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){

	//获取当前时间
	now :=time.Now()
	fmt.Println(now)

	//通过now过去 年月 日 时 分 秒
	fmt.Println(now.Year(),int(now.Month()),now.Day(),now.Hour(),now.Minute(),now.Second())

	//格式化日期时间
	nowtime :=fmt.Sprintf("%d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(nowtime)


	//格式化日期时间第二种方式
	nowtime2 := now.Format("2006-01-02 15:04:05")
	fmt.Println(nowtime2)

	//休眠  每间隔一秒钟打印一个数字，打印到100时就退出
	//i := 0
	//for  {
	//	i++
	//	fmt.Println(i)
	//	time.Sleep(time.Second)
	//	if i == 100 {
	//		break
	//	}
	//}

	//time unix 和 unixnano
	fmt.Println(now.Unix(),now.UnixNano())

	fmt.Println("---------------------------------")
	////获取改函数的执行时间
	start := time.Now().Unix()
	test3()
	end := time.Now().Unix()
	fmt.Println("执行函数的时间是",end-start,"秒")

}

func test3()  {
	str := ""
	for i:=0;i<100000; i++{
		str += "hello"+strconv.Itoa(i)
	}
}


