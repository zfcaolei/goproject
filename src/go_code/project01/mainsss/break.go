package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	var i int
	var counts int8 = 20

	for i=0;i<=int(counts);i++ {
		if i == 11 {
			break
		}
		fmt.Println(i)
	}


	fmt.Println("--------------------------------------------------")


	var count int  = 0

	for  {
		rand.Seed(time.Now().UnixNano())   //获取随机数
		n := rand.Intn(100)+1
		fmt.Println("n=",n)
		count ++
		if n == 9 {
			break
		}


	}
	fmt.Println("循环次数",count)

	fmt.Println("------------------------------------------------------")

			//break默认跳出最近的for循环
	lable1:   //lable设置为标签假设第二层循环达到条件可以中断break最外面循环
	for i:=0;i<=4;i++{
		for j:=0;j<=10;j++ {
			if j ==2 {
				break lable1
			}
			fmt.Println("j=",j)
		}
	}

	fmt.Println("------------------------------------------")


	//100以内的数求和
	lable2:
	for a:=0;a<=100;a++{
		for b:=0;b<=a;b++ {
			if a+b == 20 {
				fmt.Println("a=",a,"b=",b,"a+b=",a+b)
				break lable2
			}
		}

	}
	fmt.Println("------------------------------------------")



	//用户有三次机会输入正确的昵称及密码，如果错误就提示还剩多少次
	var user_name string
	var pass_word  string
	var cou  int = 3

	for cs:=1;cs<=cou;cs++ {
		fmt.Println("请输入用户名！！")
		fmt.Scanln(&user_name)
		fmt.Println("请输入密码！！")
		fmt.Scanln(&pass_word)

		sf := cou-cs

		if user_name == "曹磊" && pass_word == "123" {
			fmt.Println("输入正确")
			break
		}else{
			if sf == 0 {
				fmt.Println("机会已经用光")
			}else{
				fmt.Println("输入错误你还有",sf,"次机会")
			}
		}

	}
	//
	//if sf == 0 {
	//	fmt.Println("机会已经用光")
	//}


}
