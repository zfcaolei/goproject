package main

import "fmt"

type Good struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type Tv struct {
	Good
	Brand
}

type Tv2 struct {
	Good
	int //匿名字段  有一个int类型字段就不能有第二个 如果有多个就必须名字区分
}

func main() {

	//第一种写法
	tv1 := Tv{
		Good{Name: "电视", Price: 6585.99},
		Brand{Name: "索尼", Address: "日本"},
	}
	fmt.Println(tv1)

	//第二种写法
	tv2 := Tv{
		Good{"电视1", 6585.99},
		Brand{"索尼1", "日本1"},
	}
	fmt.Println(tv2)

	var tv3 Tv2
	tv3.Good.Name = "小米"
	tv3.Good.Price = 5485.55
	tv3.int = 3

	var tv4 = Tv2{
		Good{"华为", 5888},
		33,
	}
	fmt.Println(tv4)
}
