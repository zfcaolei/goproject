package main

import "fmt"

type Peopel struct {
	Age int
}

func main() {

	var peo Peopel
	peo.Age = 20
	res, res2 := peo.PeopleTest(10)
	fmt.Println(res, res2)
}

func (p Peopel) PeopleTest(a int) (int, int) {
	count := a * p.Age
	sff := a + p.Age
	return count, sff
}
