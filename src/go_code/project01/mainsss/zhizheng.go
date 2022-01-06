package main

import "fmt"

func main()  {

	var i int  = 1111
	var b int = 2222
	var rr *int
	rr  = &i
	*rr = 2000
	rr  = &b
	*rr = 3000
	fmt.Println(i,b)


	var ts *int = &i
	*ts = 222222222
	fmt.Println(*ts,i)

	//var i float32 = 0.55
	//
	//var rr *float32 = &i

	//*rr = 0.88


	 //var a int = 111
	 //
	 //var b int = 222
	 //
	 //var ccc *int
	 //
	 //ccc = &a
	 //
	 //*ccc =333
	 //
	 //ccc = &b
	 //
	 //*ccc = 444

	//var t bool = true
	//var tr *bool
	//
	//tr = &t
	//
	//*tr = false


	var str = "ccccc"

	var str_s *string = &str

	*str_s = "caolei"

	 fmt.Println(str)
	fmt.Println("-----------------")

	n1 := 10
	fmt.Println(n1)
	test41(&n1)
	fmt.Println(n1)

}

func test41(n1 *int)(n2 int)   {
	*n1 += 10
	n2 = *n1
	return
}

//func zeroval(ival int) {
//	ival = ival
//	//fmt.Println(ival)
//}
//
//func zeroptr(iptr *int) {
//	*iptr = 0
//}
//func main() {
//	i := 1
//	fmt.Println("initial:", i)
//
//	zeroval(3)
//	//fmt.Println("zeroval:", i)
//
//	// &操作符用来取得i变量的地址
//	zeroptr(&i)
//	fmt.Println("zeroptr:", i)
//	//
//	//// 指针类型也可以输出
//	//fmt.Println("pointer:", &i)

