package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	//创建一个新文件，写入内容、
	filepath := "E:/goland/testfile.txt"
	file,err :=os.OpenFile(filepath,os.O_WRONLY | os.O_CREATE,0666)

	//覆盖原先内容 file,err :=os.OpenFile(filepath,os.O_WRONLY | os.O_TRUNC,0666)
	//原本内容追加 os.O_APPEND
	if err != nil {
		fmt.Println("打开文件错误 err=",err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	str := "hello golang\n"
	writer := bufio.NewWriter(file)
	for i:=0;i<5;i++ {
		writer.WriteString(str)
	}
	writer.Flush()

	//writer是带缓存的，内容是先写入缓存的，需要调用flush方法将带缓冲的数据写入
}