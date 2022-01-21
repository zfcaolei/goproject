package main

import "fmt"

func main() {

	//接收用户的选择
	var key int

	//判断是否继续显示菜单
	//var loop  = true

	for {
		fmt.Println("-------------------------欢迎登陆多人聊天系统-------------------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:

		}
	}

}
