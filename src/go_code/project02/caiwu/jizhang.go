package caiwu

import "fmt"

type Jizhang struct {
	logout string		//结束退出确认y字段
	logintype bool		//循环字段判断是否一直for循环
	count float64		//总金额字段
	key int    			//输入状态 1=>收支明细 2=>登记收入 3=>登记支出 4=>退出软件
	ctype string		//判断支出和收入类别显示
	Mondata []Money		//存储记账数据数组
}


type Money struct {
	MoneyType int8		//记账类别  1=>收入 2=>支出
	Name string			//记账名称
	Num float64			//记账数量
}



func (m Money)Zhichu(J *Jizhang)   {
	fmt.Println("-------------------------登记支出----------------------------")
	m.MoneyType = 1
	fmt.Println("请输入支出名称！！！")
	fmt.Scanln(&m.Name)
	fmt.Println("请输入支出数量！！！")
	fmt.Scanln(&m.Num)
	J.Mondata = append(J.Mondata,m)
	J.count -= m.Num
	J.ctype = "支出"
}


func (m Money)Shouru(J *Jizhang)  {
	m.MoneyType = 2
	fmt.Println("-------------------------登记收入----------------------------")
	fmt.Println("请输入收入名称！！！")
	fmt.Scanln(&m.Name)
	fmt.Println("请输入收入数量！！！")
	fmt.Scanln(&m.Num)
	J.Mondata = append(J.Mondata,m)
	J.count += m.Num
	J.ctype = "收入"


}

func (m Money)Loginout(j *Jizhang)  {
	fmt.Println("确定要退出吗？y退出/n取消")
	fmt.Scanln(&j.logout)
	if j.logout == "y" {
		j.logintype = false
	}
}

func (m Money)Listdata(j  *Jizhang)  {
	fmt.Println("-------------------------当前收支明细-------------------------")
	if j.Mondata == nil {
		fmt.Println("-------------------------暂无信息-------------------------")
	}else {
		fmt.Println("类别----------账号金额------------收支名称--------------------金额------------")
		for _, v := range j.Mondata {
			fmt.Println(j.ctype, "        ", j.count, "           ", v.Name, "                    ", v.Num)
		}
	}
}

func (j *Jizhang)Mentsdata() {
	for j.logintype {
		fmt.Println("-----------------------家庭收支记账软件----------------------")
		fmt.Println("						1 收支明细")
		fmt.Println("						2 登记收入")
		fmt.Println("						3 登记支出")
		fmt.Println("						4 退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&j.key)
		var Mondatas Money
		switch j.key {
		case 1:
			Mondatas.Listdata(j)
		case 2:
			Mondatas.Shouru(j)
		case 3:
			Mondatas.Zhichu(j)
		case 4:
			Mondatas.Loginout(j)
		default:
			fmt.Println("请输入正确的选项！！！！")
		}
	}

}

func  NewJizhang() *Jizhang  {
	return &Jizhang{
		logintype: true,
		count : 0.00,
		logout : "",
	}
}

