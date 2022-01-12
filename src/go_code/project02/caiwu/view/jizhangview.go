package main

import (
	"fmt"
	"go_code/project02/caiwu/model"
	"go_code/project02/caiwu/service"
)

type JizhangView struct {
	logout string		//结束退出确认y字段
	logintype bool		//循环字段判断是否一直for循环
	count float64		//总金额字段
	key int    			//输入状态 1=>收支明细 2=>登记收入 3=>登记支出 4=>退出软件
	jizhangService *service.JizhangArr
}

func (this *JizhangView)Listdata()  {
	fmt.Println("-------------------------当前收支明细-------------------------")
	if this.jizhangService.JizhangData == nil {
		fmt.Println("-------------------------暂无信息-------------------------")
	}else {
		fmt.Println("收支类别--------------总金额------------------收支名称--------------------金额------------")
		for K, _ := range this.jizhangService.JizhangData {
			this.jizhangService.JizhangData[K].GetInfo(this.count)
		}
	}
}

func (this *JizhangView)AddView(typeMoney int8)  {
	fmt.Println("-------------------------登记收入----------------------------")
	var name string
	fmt.Println("请输入收入名称！！！")
	fmt.Scanln(&name)
	var num float64
	fmt.Println("请输入收入数量！！！")
	fmt.Scanln(&num)
	if typeMoney == 1 {
		this.count += num
	}else{
		this.count -= num
	}
	moneydata := model.NewMoney(typeMoney,name,num)
	is_ok := this.jizhangService.Add(moneydata)
	if is_ok {
		fmt.Println("-------------------------登记收入成功----------------------------")
	}
}

func (this *JizhangView)Logout()  {
	fmt.Println("确定要退出吗？y退出/n取消")
	fmt.Scanln(&this.logout)
	if this.logout == "y" || this.logout == "Y" {
		this.logintype = false
		fmt.Println("-------------------------退出成功----------------------------")
	}
}

func (this *JizhangView)Title()  {
	fmt.Println("-----------------------家庭收支记账软件----------------------")
	fmt.Println("						1 收支明细")
	fmt.Println("						2 登记收入")
	fmt.Println("						3 登记支出")
	fmt.Println("						4 退出软件")
	fmt.Println("请选择(1-4):")
	fmt.Scanln(&this.key)
}


func (this *JizhangView)MenuView()  {

	zhichumag,Shoumag := 1,2

	for this.logintype {
		this.Title()
		switch this.key {
		case 1:
			this.Listdata()
		case 2:
			this.AddView(int8(zhichumag))
		case 3:
			this.AddView(int8(Shoumag))
		case 4:
			this.Logout()
		default:
			fmt.Println("请输入正确的选项！！！！")
		}
	}

}

func main()  {
	jizhangview := JizhangView{
		logintype:true,
	}
	jizhangview.jizhangService = service.NewJizhangArr()
	jizhangview.MenuView()
}