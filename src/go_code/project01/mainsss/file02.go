package main

import (
	"fmt"
	"os"
)

//文件操作 不带缓冲
func main()  {

	//一次性将文件读取到位
	file := "C:/Users/windows 10/Desktop/极验更新配置.txt"
	filedata,err := os.ReadFile(file)
	if err !=nil {
		fmt.Println("读取失败 err=",err)
	}
	fmt.Println(string(filedata))
}