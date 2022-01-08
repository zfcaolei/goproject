package main

import (
	"fmt"
)

type Utils struct {
}

func main() {
	var u Utils
	data := u.UtilsList()
	fmt.Println(data)

	fmt.Println()

	res := u.Area(21.34, 45.67)
	fmt.Println(res)

	arr2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var arr3 [3][3]int

	for k1, v := range arr2 {
		fmt.Println()
		for k2, v2 := range v {
			arr3[k2][k1] = v2
		}
	}
	fmt.Println(arr3)

}

func (u Utils) UtilsList() int {
	return 10 * 8
}

func (u Utils) Area(w float64, l float64) float64 {
	return w * l
}
