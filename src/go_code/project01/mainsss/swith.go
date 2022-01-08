package main

import "fmt"

func main() {
	//var types string
	//fmt.Println("请输入性别")
	//fmt.Scanln(&types)
	//
	//switch types {
	//case "男":
	//	fmt.Println("性别：",types)
	//case "女":
	//	fmt.Println("性别：",types)
	//default:
	//	fmt.Println("性别未知")
	//
	//}

	var n1 int16 = 10
	//var n2 int16 = 20

	//可以携带多个表达式   default 语句不是必须的
	//switch n1 {
	//case n2,40:
	//	fmt.Println("n2")
	//case 30:
	//	fmt.Println("30")
	//default:
	//	fmt.Println("------")
	//}

	//switch 后也可以不带表达式 类型if else

	//switch  {
	//case n1 == 10:
	//	fmt.Println("---10--",n1)
	//case n1 == 20:
	//	fmt.Println("--20---",n1)
	//default:
	//	fmt.Println("--***---",n1)
	//}

	//switch 的穿透 fallthrougth
	switch n1 {
	case 10:
		fmt.Println("---10-----------", n1)
		fallthrough
	case 20:
		fmt.Println("---20--------------", n1)
	default:
		fmt.Println("--***---", n1)
	}

}
