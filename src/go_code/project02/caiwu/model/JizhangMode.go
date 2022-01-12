package model

import "fmt"

type Money struct {
	MoneyType int8		//记账类别  1=>收入 2=>支出
	Name string			//记账名称
	Num float64			//记账数量
}

func NewMoney(moneytype int8,name string,num float64)  Money{
	return Money{
		MoneyType:moneytype,
		Name: name,
		Num: num,
	}
}

func (this *Money)GetInfo(count float64)  {
	var MontypeName string
	if this.MoneyType == 1 {
		MontypeName = "收入"
	}else{
		MontypeName = "支出"
	}
	fmt.Println( MontypeName,"                     ",count,"                  ",this.Name, "                    ", this.Num)
}