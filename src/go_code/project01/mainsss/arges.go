package main

import (
	"fmt"
	"os"
)

//命令行参数获取
func main()  {
	fmt.Println("命令行参数有",len(os.Args))
	for k,v:=range os.Args{
		fmt.Println(k,v)
	}
}
