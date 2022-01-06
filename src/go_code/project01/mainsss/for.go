package main

import "fmt"

func main()  {

	var num int8 = 20
	var i int8

	for i=1;i<= num ;i++ {
		fmt.Println("111")
	}


	////for循环第二种写法
	//var a int8= 1
	//for a<=20 {
	//	fmt.Println(a)
	//	a++
	//}
	//
	////for循环第三种写法
	//var af int8= 1
	//for {
	//	if af<=10 {
	//		fmt.Println(af)
	//	}else{
	//		break
	//	}
	//	af++
	//
	//}

	//遍历字符串
	var str string= "caolei"
	////如果字符串遍历含有中文。那么传统遍历字符串方式就是错误会出现乱码，原因是传统的
	//	//对字符串的遍历是按照字节来遍历，而一个汉字再utf8代表3个字节
	//for i:=0;i<len(str);i++ {
	//	fmt.Printf("%c \n",str[i])
	//}

	//第二种for range 遍历
	str = "hello,word!!!!"
	for key,val := range str{
		fmt.Printf("key= %d, val=%c \n",key,val)
		fmt.Println(key,val)
	}

	//m := map[string]int{
	//	"hello": 100,
	//	"world": 200,
	//}
	//for key, value := range m {
	//	fmt.Println(key, value)
	//}

	var count int64 = 0
	var sum int64 = 0
	var as int64 = 0
	for ;as<=100;as++ {
		if as % 9 == 0 {
			sum += as
			count++
		}
	}
	fmt.Println("总和=",sum)
	fmt.Println("总个数=",count)

	var is int = 12121212
	fmt.Println(is)
	fmt.Println("-----------------------------")
	var iss int = 0
	for  {
		if iss > 20 {
			break
		}
		fmt.Println(iss)
		iss ++
	}

	fmt.Println("-----------------------------")
	var a int
	for a=0;a<=10;a++ {
		fmt.Println("循环体",a)
	}
	fmt.Println("循环外",a)

}