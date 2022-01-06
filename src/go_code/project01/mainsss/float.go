package main

import (
	"fmt"
	//"github.com/shopspring/decimal"
)

func main() {

	//浮点位数部分可能会丢失
	var f1 float32 = -123.0000901
	var f2 float64 = 123.000000901
	fmt.Println(f1, f2)

	//浮点默认声明是float64类型
	var f3 = 125.58

	f4 := .545 // = 0.545
	fmt.Println(f3, f4)


	var price float64 = 60
	fmt.Println(price/3)


	 var f11 float64 = 0.3
	 var f22 float64 = 0.6
	 fmt.Println(f11 + f22)


}
