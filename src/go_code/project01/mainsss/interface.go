package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}


//手机
type Phont struct {

}

func (p Phont)Start()  {
	fmt.Println("手机开始工作")
}

func (p Phont)Stop()  {
	fmt.Println("手机结束工作")
}

//相机-
type  Camert struct {

}

func (c Camert)Start()  {
	fmt.Println("相机开始工作")
}

func (c Camert)Stop()  {
	fmt.Println("相机结束工作")
}


//电脑
type Compoer struct {
	
}

func (compoer Compoer)Work(usb Usb)  {
	usb.Start()
	usb.Stop()
}

func main()  {
	compoer := Compoer{}
	phone := Phont{}
	camert :=Camert{}

	compoer.Work(phone)
	compoer.Work(camert)
}