package service

import "go_code/project02/caiwu/model"

type JizhangArr struct {
	JizhangData []model.Money
}

func NewJizhangArr()*JizhangArr  {
	return &JizhangArr{}
}

func (this *JizhangArr)Add(money model.Money)bool  {
	this.JizhangData = append(this.JizhangData,money)
	return true
}

