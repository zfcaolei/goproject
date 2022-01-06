package main

import "fmt"
var name string = "caolei"

func sum(n1,n2 int)int  {
	fmt.Printf("n1 type=%T\n",n1)
	return  n1+n2
}

func test5()  {
	fmt.Println(name)
}

func main()  {
	res :=sum(1,2)
	fmt.Println(res)


	fmt.Println("-----------")
	n2,n1 :=swap(1,2)
	fmt.Println(n1,n2)


	fmt.Println("---")
	a :=1
	b :=2
	swaps(&a,&b)
	fmt.Println(a,b)


	c :=1
	d :=2

	res = c
	c = d
	d = res
	fmt.Println("c=",c,"d=",d)


	fmt.Println("--------------------------------------")

	test5()


}

func swap(n1 int,n2 int)(int,int)  {
	//n3 = n2
	//n4 = n1
	return n1,n2
}

func swaps (n1 *int,n2 *int){
	res := *n1
	*n1 = *n2
	*n2 = res
}