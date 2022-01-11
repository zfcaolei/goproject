package main

import "fmt"

type AAinterface interface {
	say()
}

type BBinterface interface {
	run()
}

type Monter struct {

}

func (m Monter)say()  {
	fmt.Println("AAinterface----say()")
}

func (m Monter)run()  {
	fmt.Println("BBinterface----run()")
}



func main()  {

	//实现A B 接口
	var m Monter

	var aai AAinterface  =  m
	var bbi BBinterface  =  m
	aai.say()
	bbi.run()
}