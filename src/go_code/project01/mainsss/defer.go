package main

import "fmt"

//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈
//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
func sums(n1 int,n2 int) int{
	defer fmt.Println("1=",n1)
	defer fmt.Println("2=",n2)

	res := n1+n2
	fmt.Println("3=",res)
	return  res
}


func sums2(n1 int,n2 int) (res int){
	defer fmt.Println("1=",n1) //defer语句把n1的时候就入栈了就把值带入进去了 后面有累加也不会变
	defer fmt.Println("2=",n2)
	n1 ++
	n2 ++

	res = n1+n2
	fmt.Println("3=",res)
	return
}


func c() (i int) {
	defer func() { i++ }()
	return 1
}

// return 之后修改函数的返回值.
func doubleSum(a, b int) (sum int) {
	defer func() {   //该函数在函数返回时  调用
		sum *= 2
	}()
	sum = a + b
	return  sum
}


func main(){

	res :=c()
	fmt.Println(res)
	fmt.Println("---------------------")

	res2:=sums2(1,2)
	fmt.Println(res2)

	fmt.Println("---------------------")


	res3 := doubleSum(1,2)
	fmt.Println(res3)
}