package main

import "fmt"

//init 比main函数先执行 通常再init函数中完成初始化操作
func init()  {
	fmt.Println("init")
}

func main()  {
	fmt.Println("main")
}