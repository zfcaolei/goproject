package main

import (
	"fmt"
	"sort"
)
//对结构体进行排序
type Peoples struct {
	Name string
	Age int
	Class int8
}

type PeopleList []Peoples
type  sad string

func (p PeopleList) Len() int{
	return len(p)
}
func (p PeopleList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PeopleList) Less(i, j int) bool {
	return p[i].Class > p[j].Class
}


func main()  {

	var PeopleList PeopleList
	for i:=0;i<=10;i++ {
		Peoples  := Peoples{fmt.Sprintf("曹%d",i),i,int8(i)}
		PeopleList = append(PeopleList,Peoples)
	}

	sort.Sort(PeopleList)
	fmt.Println(PeopleList)
}