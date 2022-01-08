package main

import (
	"fmt"
	"go_code/project01/factory"
)

//工厂模式
func main() {
	student := factory.NewStudent("DSAD", 123)
	fmt.Println(*student)
}
