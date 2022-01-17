package main

import "fmt"

//type slicearr [][]int

func rrssss(arr *map[int]int,bum int)  {
	(*arr)[bum] = bum
}

func tets(channel_arr []chan bool,i int)  {
	channel_arr[i] = make(chan bool,1)
	channel_arr[i] <- true
	close(channel_arr[i])
}

func main()  {

	//var ss interface{}
	//ss = "sadsa"
	//
	//
	////if ss, ok := ss.(string), ok {
	////	//正确的情况
	////}{
	////	//错误的情况
	////}
	//
	//if valueInt, ok := ss.(string);ok{
	//	fmt.Println(1111,valueInt)
	//}else{
	//	fmt.Println(2222,valueInt)
	//}
	//for i:=1;i<=6;i++ {
	//	fmt.Println(i)
	//
	//}

	dataarr := make(chan map[string]string,200)
	  rs := make(map[string]string,200)
	bs := make(map[string]string,200)
	rs["name"] = "曹磊"
	dataarr <- rs

	bs["age"]  = "18"
	dataarr <- rs

	ma := <-dataarr
	fmt.Println(ma)


	sda := make(map[int]int)
	rrssss(&sda,1)
	rrssss(&sda,2)
	rrssss(&sda,3)
	fmt.Println(sda)

	exitchan := make(chan int)
	_,ok := <-exitchan
	if ok {
		fmt.Println(ok)
	}

}
