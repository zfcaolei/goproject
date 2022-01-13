package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件操作
func main()  {

	//打开文件
	file,err := os.Open("C:/Users/windows 10/Desktop/极验更新配置.txt")
	if err != nil {
		fmt.Println("open file err=",err)
	}
	//当函数退出是要及时关闭
	defer file.Close()

	read := bufio.NewReader(file)

	//循环的读取文件的内容
	for  {
		//带缓冲的读取每次只读一些适合于大文件
		str,err := read.ReadString('\n')
		if err == io.EOF {
			 break
		}
		fmt.Println(str)
	}
	fmt.Println("读取完成")
}