package main

import "fmt"

type tes struct {
	Code int
	Msg  string
	Data struct {
			OrderNo     string
			PayAmount   string
	}
}

type Data struct {
	OrderNo     string
	PayAmount   string
}

func main()  {
	//data := Data{OrderNo: "#21",PayAmount: "3213"}

	stes := tes{Code: 200,Msg: "失败",Data: Data{OrderNo: "#21",PayAmount: "3213"}}

	fmt.Println(stes)

}