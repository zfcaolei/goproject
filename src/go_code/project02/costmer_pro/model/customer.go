package model

import "fmt"

type Costomer struct {
	Id int			//编号id
	Name string		//年龄
	Gender string  //性别
	Age int		   //年龄
	Phont string   //手机
	Email string  //邮箱
}

//func NewCostomer(name string,gender string,age int,phone string,email string)Costomer {
//	return Costomer{
//			Name: name,
//			Gender: gender,
//			Age: age,
//			Phont: phone,
//			Email: email,
//	}
//}

func (coster *Costomer)Getinfo()string  {
	info :=fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",coster.Id,coster.Name,coster.Gender,coster.Age,coster.Phont,coster.Email)
	return  info
}



