package main

import "fmt"

func main()  {

	var t1 = [3]int{11,22,33}
	var s1 = t1[:]
	s1[1] = 11111111111111

	fmt.Println(s1)
	fmt.Println(t1)

}

func test(m []int)  {
	m[1] = 22222222222222
}

