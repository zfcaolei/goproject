package main

import "fmt"

//抽象
type Bank struct {
	Account string
	Pwd     int
	Num     float64
}

//查询余额
func (bank *Bank) Select(pwd int) {
	if bank.Pwd != pwd {
		fmt.Println("密码错误")
		return
	}
	fmt.Println("您的账号余额:", bank.Num)
}

//存款
func (bank *Bank) Insert(pwd int, num float64) {
	if bank.Pwd != pwd {
		fmt.Println("密码错误")
		return
	}
	bank.Num += num
	fmt.Println("存款成功！")
}

//取款
func (bank *Bank) Qukuang(pwd int, num float64) {
	if bank.Pwd != pwd {
		fmt.Println("密码错误")
		return
	}
	if bank.Num < num {
		fmt.Println("金额不足,取款失败！")
		return
	}
	bank.Num -= num
	fmt.Println("取款成功！")
}

func main() {
	var BankData = Bank{
		Account: "gs123456",
		Pwd:     88888888,
		Num:     10.0,
	}
	for true {
		var types int
		var pwd int
		var num float64
		fmt.Println("1=>查询 2=>取款 3=>存款,请输入操作码！！！")
		fmt.Scanln(&types)
		fmt.Println("请输入您的卡号密码！")
		fmt.Scanln(&pwd)
		switch types {
		case 1:
			BankData.Select(pwd)
		case 2:
			fmt.Println("请输入取款金额！")
			fmt.Scanln(&num)
			BankData.Qukuang(pwd, num)
		case 3:
			fmt.Println("请输入存款金额！")
			fmt.Scanln(&num)
			BankData.Insert(pwd, num)

		}
	}
	//BankData.Select(88888888)
	//BankData.Insert(88888888,20)
	//BankData.Select(88888888)
	//BankData.Insert(88888888,50)
	//BankData.Select(88888888)
	//BankData.Qukuang(88888888,30)
	//BankData.Select(88888888)

}
