package main

import (
	"fmt"
)

//基本数据类型转换
func main() {
	var i int16 = 9999

	var s float32 = float32(i)

	var a int32 = int32(i) //低精度到高精度

	//大数据类型转小数据类型编译不会报错，只是装换按溢出处理
	var g int8 = int8(i)

	//运算需要装换当前类型
	var cc int32 = 12
	var vv int8 

	vv = int8(i) + 127

	fmt.Println(i, s, a, g, cc, vv)
}
