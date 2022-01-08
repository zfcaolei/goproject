package main

import "fmt"

type AA struct {
	Name string
	age  int
}

type BB struct {
	Name  string
	Score float64
}

type CC struct {
	AA
	BB
}

type DD struct {
	a AA //有名结构体   组合关系
}

func main() {

	//如果C没有name字段，而A和B有，这时就必须通过指定匿名结构体名字来区分
	var cc CC
	cc.AA.Name = "CCC"
	cc.BB.Name = "aaaa"
	fmt.Println(cc)

	var dd DD
	dd.a.Name = "发个"
	fmt.Println(dd)
}
