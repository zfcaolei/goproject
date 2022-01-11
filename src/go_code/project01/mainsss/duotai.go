package main

import "fmt"

type Usbz interface {
	Start()
	Stop()
}


//手机
type Phontz struct {
	Name string
}

func (p Phontz)Start()  {
	fmt.Println("手机开始工作")
}

func (p Phontz)Stop()  {
	fmt.Println("手机结束工作")
}

//相机-
type  Camertz struct {
	Name string
}

func (c Camertz)Start()  {
	fmt.Println("相机开始工作")
}

func (c Camertz)Stop()  {
	fmt.Println("相机结束工作")
}


//电脑
type Compoerz struct {

}

func (compoer Compoerz)Works(usb Usbz)  {
	usb.Start()
	usb.Stop()
}


func main()  {
	//compoer := Compoerz{}
	//phone := Phontz{}
	//camert :=Camertz{}
	//
	//compoer.Works(phone)
	//compoer.Works(camert)
	var sad  [3]Usbz
	sad[0] = Phontz{"小米"}
	sad[1] = Phontz{"苹果"}
	sad[2] = Camertz{"索尼"}
	fmt.Println(sad)

}