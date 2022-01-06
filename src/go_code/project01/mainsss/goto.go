package main

import "fmt"

func main()  {

	var types bool = true

	fmt.Println("ok1")
	if types {
		goto lable4
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	lable4:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")





}