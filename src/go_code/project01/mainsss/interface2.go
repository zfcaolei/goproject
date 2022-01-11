package main

import "fmt"

/*接口*/
type Ainterface interface {
	say()
}



/*学生*/
type Students struct {
	name string
}

func (s Students)say()  {
	fmt.Println("学生调用say")
}



/*医生*/
type Doctor struct {
	name string
}

func (d Doctor) say()  {
	fmt.Println("医生调用say")
}


/*自定义类型*/
type inter int

func (i inter)say()  {
	fmt.Println("inter调用say")
}


func main()  {

	var stu Students
	var stuainderface = stu
	stuainderface.say()


	var doctor Doctor
	var doctorinterface = doctor
	doctorinterface.say()


	var i inter
	var iinterface = i
	iinterface.say()

}
