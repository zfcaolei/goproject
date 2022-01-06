package main

import "fmt"

//面向对象


type Cat struct {
	Name string
	Age int8
	Color string
	Dongzuo map[string]string //映射
	Dongzuo2 map[string]map[string]string //二维映射
	Slicedata []int //切片
	Arrdata [2]int //数组
	Arrdata2 [2][2]string //二维数组
}

func main()  {
	//方式1
	var cat1 Cat
	cat1.Name = "花花"
	cat1.Age = 11
	cat1.Color = "白色"
	cat1.Dongzuo = make(map[string]string)
	cat1.Dongzuo["name"] = "曹磊"
	cat1.Dongzuo["age"] = "15"

	cat1.Dongzuo2 = make(map[string]map[string]string)
	cat1.Dongzuo2["bo1"] = make(map[string]string)
	cat1.Dongzuo2["bo1"]["name"] = "carl"
	cat1.Dongzuo2["bo1"]["sex"] = "男"
	cat1.Dongzuo2["bo2"] = make(map[string]string)
	cat1.Dongzuo2["bo2"]["name"] = "marl"
	cat1.Dongzuo2["bo2"]["sex"] = "女"


	cat1.Slicedata = make([]int,2)
	cat1.Slicedata[1] = 1
	cat1.Slicedata = []int{0:1,2:3,3:4}


	cat1.Arrdata[0] = 11
	cat1.Arrdata[1] = 12


	cat1.Arrdata2[0][0] = "das"
	cat1.Arrdata2[0][1] = "das"
	cat1.Arrdata2[1][0] = "dsa"
	cat1.Arrdata2[1][1] = "as"
	fmt.Println(cat1)
	fmt.Println(cat1.Color)

}