package main

import (
	"fmt"
)

//go数组

func main()  {



	//定义数组

	var  hens [6]float64

	hens[0] = 10.0
	hens[1] = 10.1
	hens[2] = 10.2
	hens[3] = 10.3
	hens[4] = 10.4
	total := 0.0
	for _,val :=range hens{
		total +=val
	}
	fmt.Println(total)
	fmt.Println("-------------------------------------")

	//其他四种初始化数组方式
	var test1 [3]int = [3]int{4,5,6}
	fmt.Println(test1)

	var test5 = [3]string{"222","ck","申诉"}
	fmt.Println(test5)

	//...固定写法
	var test6 = [...]int{23,3,7,54,7}
	fmt.Println(test6)

	//指定下标
	var test7 = [...]int{1:800,0:32,4:7,3:34,2:7}
	fmt.Println(test7)

	fmt.Println("-------------------------------------")
	var test [4] int

	for key,_ := range test{
		test[key] = 20
	}

	fmt.Println(test)

	fmt.Println("-------------------------------------------")

	//函数引用数组
	arrs := [...]int{12,34,56}
	test33(&arrs)
	fmt.Println(arrs);

	fmt.Println("-------------------------------------------")

	var count [4]int


	for key,_ := range count{
		fmt.Println("请输入数组值！！！")
		fmt.Scanln(&count[key])
	}

	fmt.Println(count)




}

func test3s(arr [3]int)  {
	arr[0] = 11
}


func test33(arr *[3]int)  {
	(*arr)[0] = 11
}