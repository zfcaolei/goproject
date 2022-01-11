package main

import "fmt"

//type slicearr [][]int


func main()  {

	var ss interface{}
	ss = "sadsa"


	//if ss, ok := ss.(string), ok {
	//	//正确的情况
	//}{
	//	//错误的情况
	//}

	if valueInt, ok := ss.(string);ok{
		fmt.Println(1111,valueInt)
	}else{
		fmt.Println(2222,valueInt)
	}

}
