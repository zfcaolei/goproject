package main

import (
	"fmt"
	"go_code/project01/test"
	"go_code/project01/utils"
	"go_code/project01/while"
	bus "go_code/project02/buyt"  //可以给包取别名
)

func main()  {

	var tests int = while.Wt(3232)
	fmt.Println(tests)

	var ns  = utils.Cl("dasdas")
	fmt.Println(ns)

	var testname string = testssss.Tests("我的")
	fmt.Println(testname)

	 hs,_ := bus.Str(3,2)
	fmt.Println(hs)
	 as,xs:= bus.Str(3,2)
	fmt.Println(as,xs)

	 fmt.Println(bus.Str(3,2))


	fmt.Println("buyt包里面的变量Num1=",bus.Num)
}