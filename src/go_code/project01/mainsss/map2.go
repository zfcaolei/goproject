package main

import "fmt"

//map
//一维

func main()  {
	m1 := map[string]string{
		"name":"ccc",
		"age" : "323",
	}
	fmt.Println(m1)


	var mm1 = make(map[string]string)
	mm1["name"]	 = "cccc1"
	mm1["age"] = "2323"
	fmt.Println(mm1)

	//关联
	m2 := map[string]map[string]string{
		"ob1" :{"name":"ccccc2","age":"222222"},
		"ob2" :{"name":"ccccc3","age":"3333333"},
	}
	fmt.Println(m2)


	var mm2 = make(map[string]map[string]string)
	mm2["ob1"] = make(map[string]string)
	mm2["ob1"]["name"] = "ccccc3"
	mm2["ob1"]["age"]  = "33333"
	mm2["ob2"] = make(map[string]string)
	mm2["ob2"]["name"] = "ccccc555"
	mm2["ob2"]["age"]  = "666666"
	fmt.Println(mm2)
}
