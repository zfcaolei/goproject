package main

import "fmt"



//大猴子
type Monkey struct {
	Name string
}

//大猴子技能
func (this *Monkey)climbing() {
	fmt.Println(this.Name,"生来会爬树")
}

//小猴子继承大猴子
type SmallMonkey struct {
	Monkey
}


//鸟类
//接口飞翔
type Bird interface {
	Fly()
}


func (this *SmallMonkey) Fly() {
	fmt.Println(this.Name,"通过学习会飞翔")
}





func main()  {
	
	var smallmonkey = SmallMonkey{Monkey{"小猴子"}}
	smallmonkey.climbing()
	smallmonkey.Fly()
}
