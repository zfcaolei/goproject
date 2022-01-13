package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件/图片/视频拷贝

func CopyFile(dstFilename string,srcFilename string)(writen int64,err error)  {

	srcfile,err := os.Open(srcFilename)
	if err!=nil {
		fmt.Println("open file err=",err)
	}
	//通过srcfile,获取到Reader
	reader := bufio.NewReader(srcfile)
	defer srcfile.Close()

	//打开dstfileName
	dsrFile,err :=os.OpenFile(dstFilename,os.O_WRONLY | os.O_CREATE,0666)
	if err !=nil {
		fmt.Printf("open file err=",err)
	}
	defer dsrFile.Close()
	writer := bufio.NewWriter(dsrFile)
	return io.Copy(writer,reader)
}


func main(){

	file_path := "C:/Users/windows 10/Pictures/photo_2021-02-03_17-46-23.jpg"
	new_file_path := "E:/goland/testfile.jpg"
	CopyFile(new_file_path,file_path)
}
