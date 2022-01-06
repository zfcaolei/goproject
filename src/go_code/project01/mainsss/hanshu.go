package main

import "fmt"

func main()  {
	result := results(111,333333)
	fmt.Println(result)
	test(4)
	test2(4)
	num :=20
	test4(&num)
	fmt.Println("指针传递num=",num)
}

func results(a int,b int) int  {
	return   a+b
}

func test4(n1 *int)  {
	*n1 = *n1+10
	fmt.Println("函数内部指针传递num=",*n1)
}


// 4进入大于2进入函数 n-- 就等于n=3，    n=3
// 3再进入函数大于2 n-- 等于n=2         n=2
// 2再进入函数不大于  输出 n=2  删除掉空间然后回到上面输出n=2 删除掉空间 输出2  删除掉空间 N=3
// 递归函数先执行完再删除掉输出
func test(n int)  {
	if n>2 {
		n--
		test(n)

	}
	fmt.Println("test n=",n)
}

func test2(n int)  {
	if n>2 {
		n--
		test2(n)

	}else{
		fmt.Println("test 2 n=",n)
	}

}