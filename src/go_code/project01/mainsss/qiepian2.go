package main

import "fmt"

func main()  {


	//切片
	//一维
	a1 := []int{1,2,3,4,5}
	fmt.Println(a1)

	var a2 = make([]int,5)
	a2[0] = 1
	a2[1] = 2
	a2[2] = 3
	a2[3] = 4
	a2[4] = 5
	fmt.Println(a2)

	//二维
	aa1 := [][]int{{1,2,3,4,5},{5,6,7,8,9}}
	fmt.Println(aa1)

	var aa2 = make([][]int,2)
	aa2[0] = make([]int,5)
	aa2[0][0] = 1
	aa2[0][1] = 2
	aa2[0][2] = 3
	aa2[1] = make([]int,5)
	aa2[1][0] = 4
	aa2[1][1] = 5
	aa2[1][2] = 6
	fmt.Println(aa2)


	//切片map
	aa3 := []map[string]string{
		{"name":"曹磊","age":"4324"},
		{"name":"曹磊","age":"4324"},
	}
	fmt.Println(aa3)

	aa4 := make([]map[string]string,2)
	aa4[0] = make(map[string]string)
	aa4[0]["name"] = "曹磊1"
	aa4[0]["age"] = "43243"
	aa4[1] = make(map[string]string)
	aa4[1]["name"] = "曹磊1"
	aa4[1]["age"] = "43243"
	fmt.Println(aa4)




}