package main

import "fmt"

type Money struct {
	MoneyType int
	Name int
	Num int
}

type MoneyData struct {
	Datas []Money

}




func main()  {

	s1 := Money{1,2,3}

	ca := MoneyData{}
	cs := append(ca.Datas,s1)
	fmt.Println(cs)
}
