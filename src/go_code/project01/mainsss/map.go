package main

import (
	"fmt"
	"sort"
)

//map映射集合 成为关联数组

func main()  {


	//map的使用方式1
	var maptest map[string] string
	//使用map前需要make分配数据空间
	maptest = make(map[string]string,10)
	maptest["no1"] = "曹磊"
	maptest["no2"] = "王鑫"
	maptest["no1"] = "曹磊222"
	fmt.Println(maptest)


	fmt.Println("------------------------------------------------")
	//map的使用方式2
	var test2 = make(map[string]string)
	test2["name"]= "caolei"
	test2["age"] = "19"
	fmt.Println(test2)

	var test22 = make(map[string]map[string]string)
	test22["ob1"],test22["ob2"] = make(map[string]string),make(map[string]string)
	test22["ob1"]["name"]= "caolei"
	test22["ob1"]["sex"] = "男"
	test22["ob2"]["name"] = "carl"
	test22["ob2"]["sex"] = "男"
	fmt.Println(test22)


	fmt.Println("------------------------------------------------")
	//map的使用方式3
	test3 := map[string]string{
		"he1":"大萨达",
		"he2":"广告费",
	}
	fmt.Println(test3)

	test4 := map[string]map[string]string{
		"no1":{"name":"carl","age":"22","sex":"男"},
		"no2":{"name":"mark","age":"32","sex":"男"},
	}
	fmt.Println(test4)



	fmt.Println("------------------------------------------------")
	//map的增删改查操作

	//增加
	test333 := make(map[string]string)
	test333["i1"] = "大萨"
	test333["i2"] = "sad"

	//修改
	test333["i1"] = "士大夫"

	//删除
	delete(test333,"i2")
	fmt.Println(test333)
	//删除map所有key第一种变量所有key删除
	//直接make一个新空间
	test333 = make(map[string]string)


	//查询
	val,mapres := test22["ob1"]["name"]
	if mapres {
		fmt.Println("有存在",val)
	}else{
		fmt.Println("没有不存在")
	}

	val,mapresss := test333["i2"]
	if mapresss {
		fmt.Println("有存在",val)
	}else{
		fmt.Println("没有不存在1",val)
	}


	if val, ok := test22["ob1"]["name"]; ok {
		fmt.Println("有存在",val)
	}else{
		fmt.Println("没有，不存在",val)
	}


	fmt.Println("------------------------------------------------")
	//map变量
	for _,val := range test22{
			fmt.Println(val["name"])

	}

	//map长度查询
	fmt.Println("map长度test22是",len(test22))



	fmt.Println("------------------------------------------------")
	//map切片
//	var monster []map[string]string
	monster := make([]map[string]string,2)
	if monster[0] == nil {
		monster[0] = make(map[string]string,2)
		monster[0]["name"] = "宝马"
		monster[0]["price"] = "1100000"
	}
	if monster[1] == nil {
		monster[1] = make(map[string]string,2)
		monster[1]["name"] = "奔驰"
		monster[1]["price"] = "100000"
	}


	newmonster := map[string]string{
		"name":"奥迪",
		"price":"200000",
	}

	monster = append(monster,newmonster)
	fmt.Println(monster)





	fmt.Println("------------------------------------------------")
	//map排序
	mates := make(map[int]int,10)
	mates[0] = 1
	mates[4] = 22
	mates[2] = 33
	fmt.Println(mates)
	//先将map的key放入到切片中
	//对切片排序
	//遍历切片，然后按照key来输出map的值
	//新版的golang对map已经有了排序

	var paixu  []int
	for key,_ := range mates{
		paixu = append(paixu,key)
	}
	sort.Ints(paixu)

	for _,va :=range paixu{
		fmt.Println(va,mates[va])
	}
	fmt.Println(paixu)


}


