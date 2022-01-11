package main

import "fmt"

type TwoItf interface {
	test1()
}

type ThreeItf interface {
	test2()
}

//继承以上两个接口
type OneItf interface {
	TwoItf
	ThreeItf
	test3()
}


type Stu struct {

}


func (s Stu)test3()  {
	fmt.Println("test3")
}
func (s Stu)test2()  {
	fmt.Println("test2")
}
func (s Stu)test1()  {
	fmt.Println("test1")
}

func main()  {

	var s Stu
	//使用OneItf接口,需要实现OneItf继承的所有接口的方法
	var oneitf OneItf = s

	oneitf.test1()


}