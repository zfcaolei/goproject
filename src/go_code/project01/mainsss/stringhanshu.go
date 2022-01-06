package main

import (
	"fmt"
	"strconv"
	"strings"
)

//字符串函数使用

func main()  {

	//统计字符串长度，按照字节返回
	//golang的编码统一为utf-8 字母和 数字占1个字节，汉字占用3个字节
	str1 := "hello是"
	fmt.Println(len(str1))


	str2 := "hellw背景"

	//字符串遍历有中文问题
	r := []rune(str2)

	for i:=0;i<len(r);i++ {
		fmt.Printf("字符=%c\n",r[i])
	}


	//字符串转换整数
	n,err := strconv.Atoi("123")
	fmt.Println(n,err)
	if err != nil {
		fmt.Println("装换错误")
	}else{
		fmt.Println(n)
	}


	//查找字符串中是否存在指定的字符串 返回值bool
	s := strings.Contains("fsdff","fss")
	fmt.Println(s)

	//截取字符串最后XX位字符
	b := "sda323.jpg"
	c := b[1 : len(b)-4]
	fmt.Println(c)

	//返回字符串最后一次出现的位置
	gg := strings.LastIndex("caolei","o")
	fmt.Println(gg)

	//将指定字符串替换成另外的字符串
	ff := strings.Replace("曹磊 go lang","曹磊","问我",-1)
	fmt.Println(ff)


	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,word",",")
	fmt.Println(strArr)

	//将字符串大小写装换
	sstr := "Golang"
	s3 := strings.ToLower(sstr)  //大写改小写
	s4 := strings.ToUpper(sstr)  //小写改大写
	fmt.Println(s3,s4)


	//将字符串左右空格去除
	fmt.Println(strings.TrimSpace(" 电风扇 "))

	//将字符串左右两边字符串去除
	fmt.Println(strings.Trim(" !fdsas "," !f"))


	//将字符串左边字符串去除
	fmt.Println("将字符串左边字符串去除",strings.TrimLeft("ffdsf","f"))


	//判断字符串是否已某个开头
	fmt.Println(strings.HasPrefix("www.sda.com","www"))

	//判断字符串是否已某个结尾
	fmt.Println(strings.HasSuffix("sda.png","png"))


}
