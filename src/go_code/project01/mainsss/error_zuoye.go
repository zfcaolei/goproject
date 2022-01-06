package main

import (
	"errors"
	"fmt"
)

func daynum()  {

	var day int
	for  {
	fmt.Println("请输入月份！！！")
	fmt.Scanln(&day)
	err :=checkmonth(day)
		if err!=nil {
			fmt.Println(err)
			continue
		}
	fmt.Println("当前输入月份=",getYearMonthToDay(2021,day))
	}
}



func checkmonth(month int)error  {
	if month>12 || month<1 {
		return errors.New("输入的月份错误！！！")
	}else{
		return nil
	}
}
// getYearMonthToDay 查询指定年份指定月份有多少天
// @params year int 指定年份
// @params month int 指定月份
func getYearMonthToDay(year int, month int) int {
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}

//类似php try cash处理
func yeys2(n1 int) bool {
	ress := true
	Try(func() {
		if n1>10 {
			ress = false
			panic("n1参数错误")
		}
		fmt.Println("3333333333")
	}, func(err interface{}) {
		fmt.Println(err)
	})
	return ress
}

func main()  {
	daynum()

	as := yeys2(2)
	fmt.Println(as)

	n1,err := test054(222)
	if err != nil {
		fmt.Println("执行错误的处理",err.Error())

	}else {
		fmt.Println("执行正确的处理，并返回函数返回值n1=",n1)
	}

}
//自定义错误异常
func test054(n1 int)(int,error)  {
	if n1>10{
		return -1,errors.New("参数错误！！！")
	}
	return  n1,nil
}

//类似php try cash处理
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

