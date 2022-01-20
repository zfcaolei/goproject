package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//socket 客户端
func main()  {

		//连接tcp服务
		conn, err := net.Dial("tcp", "192.168.2.9:8888")
		if err != nil {
			fmt.Println("cline连接错误 err=", err)
			return
		}

				//客户端可以发送单行数据，然后就退出
				reader := bufio.NewReader(os.Stdin) //os.stdin代表标准输入[终端]
		for {

				//从终端读取一行用户输入,并准备发送给服务器
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("readstring err=", err)
				}

				//如果用户输入的是exit就退出
				line = strings.Trim(line,"\r\n")
				if line=="exit" {
					break
				}

				//将line发送给服务器
				n, err := conn.Write([]byte(line))
				if err != nil {
					fmt.Println("客户端写入错误 err=", err)
				}
				fmt.Println("客服端发送了", n)
		}
}
