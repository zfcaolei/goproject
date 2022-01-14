package main

import "fmt"

type Money struct {
	MoneyType int
	Name      int
	Num       int
}

type MoneyData struct {
	Datas []Money
	sssss *Money
}

func main() {

	var s2 []Money

	//s1 := []Money{{23,23,23,}}
	s1 := Money{213, 32, 32}
	s2 = append(s2, s1)

	fmt.Println(s2)
}
