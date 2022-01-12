package service

import (
	"go_code/project02/costmer_pro/model"
)

type CostomerData struct {
	CostomerArr []model.Costomer
	CostomerNum int
}

func NewCostomerData() *CostomerData{
	costomerData := &CostomerData{}
	return costomerData
}

func (this *CostomerData)Add(costomer model.Costomer)bool {
	this.CostomerNum ++
	costomer.Id = this.CostomerNum
	this.CostomerArr = append(this.CostomerArr,costomer)
	return true
}

func (this *CostomerData)Delete(id int)bool  {
	index := this.FindById(id)
	if index == -1 {
		 return false
	}
	this.CostomerArr = append(this.CostomerArr[:index],this.CostomerArr[index+1:]...)
	return  true
}

func (this *CostomerData)FindById(id int)int  {
	index := -1
	for k,v := range this.CostomerArr{
		if v.Id == id {
			index = k
		}
	}
	return  index
}

func (this *CostomerData)Update(id int,costomer model.Costomer)bool  {
	index := this.FindById(id)
	costomerIddata := &this.CostomerArr[index]
	costomerIddata.Name = costomer.Name
	costomerIddata.Gender = costomer.Gender
	costomerIddata.Age = costomer.Age
	costomerIddata.Phont = costomer.Phont
	costomerIddata.Email = costomer.Email
	return true
}

func (this *CostomerData)Select(id int)(model.Costomer)  {
	index := this.FindById(id)
	costomerIddata := this.CostomerArr[index]
	return costomerIddata
}

