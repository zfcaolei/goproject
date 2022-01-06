package main

import "fmt"

func main()  {

	for i:=0;i<10;i++{
		if i==3 {
			continue
		}
		fmt.Println(i)
	}
fmt.Println("----------------------------------------")

	//打印1-100的奇数
	for a := 0; a <= 100;a++ {
		if a % 2 == 0 {
			continue
		}
		fmt.Println(a)
	}


fmt.Println("----------------------------------------------")
	here:
	for c:=0;c<2;c++ {
		for d:=1;d<4;d++ {
			if d==2 {
				continue here
			}
			fmt.Println("c=",c,"d=",d)
		}
	}

fmt.Println("-------------------------------")

	var shuru int
	var zhengshu int
	var fushu int

	for  {
		fmt.Println("请输入数字！！！！")
		fmt.Scanln(&shuru)

		if  shuru == 0{
			break
		}

		if shuru < 0 {
			fushu++
			continue
		}

		zhengshu ++

		}
		fmt.Println("正数次数=",zhengshu,"负数次数=",fushu)
	}




