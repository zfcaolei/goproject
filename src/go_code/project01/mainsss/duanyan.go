package main

import "fmt"

//类型断言

func main()  {

	var ss interface{}
	ss = "sadsa"

	if valueInt, ok := ss.(string);ok{
		fmt.Println(1111,valueInt)
	}else{
		fmt.Println(2222,valueInt)
	}


	sss := "dsad"

	aaa := 231

	asssss := true

	types (sss,aaa,asssss)

}


func types(item... interface{})  {
	for k,v := range item{
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数类型是bool类型,值是%v",k,v)
		case float64:
			fmt.Printf("第%v个参数类型是float64类型,值是%v",k,v)
		case int:
			fmt.Printf("第%v个参数类型是int类型,值是%v",k,v)
		case string:
			fmt.Printf("第%v个参数类型是int类型,值是%v",k,v)
		}

	}
}
