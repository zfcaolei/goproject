package main

import (
	"errors"
	"fmt"
	"time"
)

func tests()  {
	//使用defer+ recover 来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=",err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1/num2
	fmt.Println(res)
}


//错误处理

//1.在默认情况下，当发生错误后，程序就会退出
//2.如果我们希望当发生错误后，可以捕获到错误并进行处理，保证程序可以继续执行
//还可以再捕获到错误后，给管理员一个提示（邮件，短信）

//defer panic recover
func main()  {
	//即使tests函数报错也会继续执行下面代码
	tests()

	err:= test02(12)
	if err != nil {
		panic(err)
	}
	for  {
		fmt.Println("main下面的代码")
		time.Sleep(time.Second)
	}
}

//自定义错误
func test02(n1 int) (error) {
	if n1>10 {
		return  errors.New("n1参数错误")
	}
	return  nil
}
