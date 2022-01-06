package main

import "fmt"

func main(){

	//除法会保留整数去除小数
	var test int = 10 / 4
	fmt.Println(test)

	var test1 float64 = 11.0 / 4
	fmt.Println(test1)


	// a%b = a - a / b * b
	var test2 = -10%3
	fmt.Println(test2)


	var i  int = 1

	// 不能这样使用var s = i++    ++和--只能独立使用
	i++
	var s = i
	fmt.Println(s)


	var day int = 97
	var week int = day / 7
	var days int = day % 7
	fmt.Println(week,days)


	if days > 20 {
		fmt.Println(11111)
	}else {
		fmt.Println(2222222)
	}

	if day >10 && day <30 {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}

	if !(day>30)  {
		fmt.Println(11111)
	}else {
		fmt.Println(222222)
	}
}
