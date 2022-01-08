package main

import "fmt"

type Box struct {
	len    float64
	width  float64
	height float64
}

type People struct {
	Name string
	Age  int
}

func (box Box) GetVolum() float64 {
	return box.len * box.width * box.height
}

func (people *People) CheckMen() {

	for true {
		fmt.Println("请输入姓名")
		fmt.Scanln(&people.Name)
		if people.Name == "N" {
			fmt.Println("退出")
			break
		}
		fmt.Println("请输入年龄")
		fmt.Scanln(&people.Age)

		if people.Age > 18 {
			fmt.Println("姓名:", people.Name, "年龄:", people.Age, "门票:", 100)
		} else {
			fmt.Println("姓名:", people.Name, "年龄:", people.Age, "门票:", 10)
		}
	}

}

func main() {
	var people People
	people.CheckMen()

	//var box Box
	//fmt.Println("请输入长度")
	//fmt.Scanln(&box.len)
	//fmt.Println("请输入高度")
	//fmt.Scanln(&box.height)
	//fmt.Println("请输入宽度")
	//fmt.Scanln(&box.width)
	//count := box.GetVolum()
	//fmt.Println(count)
}
