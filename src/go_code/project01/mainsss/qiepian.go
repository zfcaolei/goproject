package main

import (
	"fmt"
)

//切片

func main()  {

	//引用方式切片
	var test   = [...]int{23,345,65,567,54}
	var sk = test[1:4]
	var sk2 = test[:]
	fmt.Println(test)
	fmt.Println(sk)
	fmt.Println(sk2)

	//切片可以继续切片
	var sk3 = sk2[0:]
	fmt.Println(sk3)



	fmt.Println("---------------------------------------------------")
	//通过make来创建切片
	var sli = make([]float64,3,20)  //类型，长度，容量（默认）
	sli[1] = 0.548
	fmt.Println(sli)
	fmt.Println(len(sli))
	fmt.Println(cap(sli))


	fmt.Println("---------------------------------------------------")
	//定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var sli2 = []string{"caolei","wangx","zhuyal"}
	fmt.Println(sli2)
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))


	fmt.Println("----------------------------------------------")
	//切片遍历
	for key,val := range sli2{
		fmt.Println(key)
		fmt.Println(val)
	}


	//切片二维数组
	//var test2  =  make([][]int,4)
	//for i := range test2 {
	//	test2[i] = make([]int, 5) // 每一行4列
	//}
	//fmt.Println(test2)


	fmt.Println("----------------------------------------------")
	//使用append内置函数对切片动态追加
	var sk5 = []int{100,200,300,400,500}
	sk5 = append(sk5,600)
	fmt.Println(sk5)
	//还可以通过append通过切片增加给自己
	//var sk7  []int
	sk7 := append(sk5,sk5...)
	fmt.Println(sk7)


	fmt.Println("----------------------------------------------")
	//切片的拷贝操作
	var f5 = []int{11,22,33,44,55,66,77}
	var f6 = make([]int,10)
	var f7 = make([]int,2)//如果定义的大小不够不会报错
	fmt.Println(f5)
	//copy(f6,f5)
	//copy(f7,f5)
	fmt.Println(f6)
	fmt.Println(f7)




	fmt.Println("----------------------------------------------")
	//string 和 slice  string底层是byte数组，因此string也可以进行切片处理
	str2 := "caolei"
	sk8 := str2[1:]
	fmt.Println(str2)
	fmt.Println(sk8)
	//string是不可以变的，也就是不可以通过str2[0] = "z"来改变
	//如果要改变要使用一下方式
	str_new := []byte(str2)
	str_new[0] = 'z'
	str2 = string(str_new)
	fmt.Println(str2)
	//细节，我们装换成[]byte只能处理英文和数字，不能处理中文
	//原因[]byte是按照字节来处理，而一个汉字，是三个字节，因此会出现乱码
	//解决办法是将string转换成[]rune即可，因为rune是按照字符来处理兼容汉字

	se :="曹磊的是"
	arrSe := []rune(se)
	arrSe[0] = '传'
	se = string(arrSe)
	fmt.Println(se)


	//冒泡排序实现
	var sss = []int{2,235,55,654,6}
	maopao(sss)
	fmt.Println(sss)


}

func maopao(arr []int)  {
	count := len(arr)
	for i:=0;i<=count;i++{
		for j:= count-1;j>i;j-- {
			if arr[j] < arr[j-1] {
				temp:=arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
}