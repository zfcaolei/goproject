package main

import (
	"fmt"
	"go_code/project02/costmer_pro/model"
	"go_code/project02/costmer_pro/service"
)

type CostomerView struct {
	Key int
	Loginstatus bool
	LoginOut string
	DeleteId int
	CostomerService *service.CostomerData

}

func (this *CostomerView)List()  {
	CostomerViewArr := *this.CostomerService
	fmt.Println("-------------------------客  户  列  表----------------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for k,_ := range CostomerViewArr.CostomerArr{
		fmt.Println(CostomerViewArr.CostomerArr[k].Getinfo())
	}
}

func (this *CostomerView)add() {
	fmt.Println("请输入用户姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入用户性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入用户年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入用户手机")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("请输入用户邮箱")
	email := ""
	fmt.Scanln(&email)
	costomer := model.Costomer{
		Name: name,
		Gender: gender,
		Age: age,
		Phont: phone,
		Email: email,
	}
	is_ok := this.CostomerService.Add(costomer)
	if is_ok {
		fmt.Println("---------------------------添  加  成  功----------------------------------")
	}
}

func (this *CostomerView)Delete()  {
	fmt.Println("-------------------------删  除  客  户----------------------------------")
	fmt.Println("请选择待删除用户编号(-1退出)")
	var DeleteId int
	fmt.Scanln(&DeleteId)
	if DeleteId != -1  {
		fmt.Println("确定是否删除?y/n")
		var deleout string
		fmt.Scanln(&deleout)
		if deleout =="y" || deleout == "Y" {
			is_ok := this.CostomerService.Delete(DeleteId)
			if is_ok {
				fmt.Println("---------------------------删  除  成  功----------------------------------")
			}
		}
	}
}

func (this *CostomerView)Exit()  {
	fmt.Println("----------------------您确定是否要退出,y退出/n取消！-------------------------")
	var LoginOut string
	fmt.Scanln(&LoginOut)
	if LoginOut == "y" || LoginOut == "Y" {
		this.Loginstatus = false
		fmt.Println("---------------------------退  出  成  功----------------------------------")
	}
}

func (this *CostomerView)Update()  {
	fmt.Println("-------------------------修  改  客  户----------------------------------")
	fmt.Println("请选择修改用户编号(-1退出)")
	var UpdateId int
	fmt.Scanln(&UpdateId)
	index := this.CostomerService.FindById(UpdateId)
	if index == -1 {
		fmt.Println("查找的用户编号不存在")
		return
	}
	Costomerarrdata :=this.CostomerService.Select(UpdateId)

	var name string
	fmt.Printf("请输入用户姓名(%s)",Costomerarrdata.Name)
	fmt.Scanln(&name)
	if name == "" {
		name = Costomerarrdata.Name
	}

	var gender string
	fmt.Printf("请输入用户性别(%s)",Costomerarrdata.Gender)
	fmt.Scanln(&gender)
	if gender == ""{
		gender = Costomerarrdata.Gender
	}

	var age int
	fmt.Printf("请输入用户年龄(%d)",Costomerarrdata.Age)
	fmt.Scanln(&age)
	if age == 0{
		age = Costomerarrdata.Age
	}

	var phone string
	fmt.Printf("请输入用户手机(%s)",Costomerarrdata.Phont)
	fmt.Scanln(&phone)
	if phone == ""{
		phone = Costomerarrdata.Phont
	}

	var email string
	fmt.Printf("请输入用户邮箱(%s)",Costomerarrdata.Email)
	fmt.Scanln(&email)
	if email == "" {
		email = Costomerarrdata.Email
	}

	costomer := model.Costomer{
		Name: name,
		Gender: gender,
		Age: age,
		Phont: phone,
		Email: email,
	}
	is_update_ok :=this.CostomerService.Update(UpdateId,costomer)
	if is_update_ok {
		fmt.Println("---------------------------修  改  成  功----------------------------------")
	}

}

func (this *CostomerView)CostomerMent()  {

	for  {
			fmt.Println("---------------------------客户信息管理软件----------------------------------")
			fmt.Println("                         1 添  加  客  户")
			fmt.Println("                         2 修  改  客  户")
			fmt.Println("                         3 删  除  客  户")
			fmt.Println("                         4 客  户  列  表")
			fmt.Println("                         5 退         出")
			fmt.Println("请选择1-4选项！！！")
			fmt.Scanln(&this.Key)

			switch this.Key {
				case 1:
					this.add()
				case 2:
					this.Update()
				case 3:
					this.Delete()
				case 4:
					this.List()
				case 5:
					this.Exit()
				default:
					fmt.Println("----------------------输入选项有误请重新输入-------------------------")
			}

			if this.Loginstatus == false {
				break
			}
	}
}



func main()  {
	costomerview := CostomerView{
		Loginstatus :true,
	}
	costomerview.CostomerService = service.NewCostomerData()
	costomerview.CostomerMent()
}