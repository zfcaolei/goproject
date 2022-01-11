package main

import "fmt"

func main()  {

	//数组
	var arr1 = [...]int{12,23,43,54}
	fmt.Println(arr1)

	var arr2 = [3]int{21,23,34}
	fmt.Println(arr2)

	arr11 := [2]int{}
	arr11[0] = 1
	arr11[1] = 11
	fmt.Println(arr11)

	arr22 :=[2]int{}
	arr22[0] = 11
	arr22[1] = 22
	fmt.Println()

	var arr33 [3]int
	arr33[0] = 1
	arr33[2] = 2
	fmt.Println(arr33)


	//二维数组

	var y1 [2][2]int

	y1[0][0] = 1
	y1[0][1] = 1
	y1[1][0] = 1
	y1[1][1] = 1
	fmt.Println(y1)

	y2 := [2][2]int{{1,2},{3,4}}
	fmt.Println(y2)



	//切片

	var s1 =  []int{1,2,3}
	s2 := []int{3,4,5}
	s3 := make([]int,3)
	s3[0] = 6
	s3[1] = 7
	s3[2] = 8
	fmt.Println(s1,s2,s3)


	var ss1 = [][]int{{11,22,33},{44,55,66}}
	var ss2 = make([][]int,2)
	ss2[0] = make([]int,3)
	ss2[0][0] = 77
	ss2[0][1] = 88
	ss2[0][2] = 99
	ss2[1] = make([]int,3)
	ss2[1][0] = 777
	ss2[1][1] = 888
	ss2[1][2] = 998

	fmt.Println(ss1,ss2)


	fmt.Println(ss2)

	ss33 := ss2[:][:]
	fmt.Println("ss32=",ss33)
	
	ssss := make(map[string]map[string]string)
	ssss["ob3"] = make(map[string]string)
	ssss["ob3"]["name"] = "sad"
	ssss["ob3"]["age"] = "31"
	fmt.Println(ssss)
	
	vcs := map[string]map[string]string{
		"ob1" :{"name":"xxx","age":"32"},
		"ob2" :{"name":"xxx","age":"32"},
	}
	//aaaa:= append(ssss,vcs)
	fmt.Println(vcs)








}
