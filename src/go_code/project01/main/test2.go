package main

import "fmt"

func main() {

	 name := "caolei" //s
	 age := 111		//d
	 sex := false  //t
	 arr := []int{12312,31231}
	arrs := [2][2]string{}
	arrs[0][0] = "萨达"
	arrs[0][1]  = "萨达"
	arrs[1][0]  = "萨达"
	arrs[1][1]  = "萨达"



	arrqiepian :=  arrs[:][:]
	arrqiepian2 := make([][]string,2)
	arrqiepian2[1],arrqiepian2[0] = make([]string,2),make([]string,2)
	arrqiepian2[0][0] = "萨达1"
	arrqiepian2[0][1]  = "萨达1"
	arrqiepian2[1][0] = "萨达1"
	arrqiepian2[1][1]  = "萨达1"


	ss := [][]string{{"3水电费","3水电费","3水电费"},{"3水电费","3水电费","3水电费"}}
	fmt.Println(ss)

	fmt.Println(arrqiepian2)
	fmt.Printf(
		"name=%s,age=%d,sex=%t,arr=%v,arrs=%v,arrqiepian=%v",
		name,age,sex,arr,arrs,arrqiepian)

}

