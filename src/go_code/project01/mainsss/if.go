package main

import "fmt"

func main(){
	//var  i = 30
	//var  d = 20
	//if i<d {
	//	fmt.Println("i小于",d)
	//}else{
	//	if i == d {
	//		fmt.Println("i等于",d)
	//	}else{
	//		fmt.Println("i大于",d)
	//	}
	//}

	//var age int
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//
	////var ages = age  也可以如下使用
	//
	//if ages:= age;ages >= 18 {
	//	if ages == 18 {
	//		fmt.Println("你今年",age,"岁等待马上成年了")
	//	}else {
	//		fmt.Println("你今年",age,"岁成年了")
	//	}
	//
	//}else{
	//	fmt.Println("你今年",age,"岁还未成年")
	//}
	//
	//
	////多分支控制
	//if ages:= age;ages == 19 {
	//	fmt.Println("你今年",age,"岁---19")
	//}else if ages == 20{
	//	fmt.Println("你今年",age,"岁---20")
	//}else if ages == 30 {
	//	fmt.Println("你今年",age,"岁---30")
	//}else {
	//	fmt.Println("你今年",age,"岁---end")
	//}


	var shenggao uint8
	var money  int32
	var look  string

	fmt.Println("请输入身高")
	fmt.Scanln(&shenggao)

	fmt.Println("请输入你的银行余额")
	fmt.Scanln(&money)

	fmt.Println("你帅吗？请输入帅和不帅！！！！！")
	fmt.Scanln(&look)

	//var looks string
	//if look == true{
	//	looks = "帅"
	//}else{
	//	looks = "丑"
	//}

	fmt.Println(" 身高：",shenggao,"cm\n","银行余额:",money,"\n","长相：",look)

	if shenggao >=180 && money >=10000000 && look == "帅" {
			fmt.Println("我一定要嫁给他")
	}else if shenggao>=180 || money>=10000000 || look == "帅"{
		  fmt.Println("嫁吧,还可以")
	}else {
		fmt.Println("不嫁")
	}

	//var age int = 20


}
