package main

import (
	"fmt"
	"runtime"
)
//go1.8版本后默认让程序运行再多核上，可以不用设置
func main()  {
	//运行cpu数量
	num := runtime.NumCPU()
	fmt.Println(num)

	//可以自己设置使用多少个cpu
	runtime.GOMAXPROCS(num-1)

}
