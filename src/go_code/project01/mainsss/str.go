package main

import (
	"fmt"
)

func main() {

	var s1 byte = 'a'
	var s2 byte = '0'
	//直接输出byte值是输出对应的字符码值
	fmt.Println(s1, s2)

	//如果想输出对应字符，需要使用格式化输出
	fmt.Printf("%c %c", s1, s2)

	var s4 int = '被'
	fmt.Printf("\nc3=%c c3对应码值=%d", s4, s4)

	//输出字符带很多符号可以用反引号
	var s5 = `dsfdsgflhk发个抵抗力换个地方来看还记得肺力咳合剂3=======`

	fmt.Println(s5)

	var strings = "hello" + "ssss"

	strings += "word"

	//拼接字符串加号需要在上面
	var sss = "fff" +
		"cccccc"

	var b bool    //默认值false
	var s string  //默认值 ""
	var a float32 //默认 0
	var c int     //默认 0
	fmt.Println(strings, sss, b, s, a, c)



	var basj byte = 'f'
	fmt.Println("basj=",basj)



}
