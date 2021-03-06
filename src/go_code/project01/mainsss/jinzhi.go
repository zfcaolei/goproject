package main

import "fmt"

func main()  {

	var i int = 5
	//二进制输出  10开头2进制
	fmt.Printf("%b \n",i)

	//八进制 0-7 满8进1 以数字0开头
	var c int = 011  //011 = 9
	fmt.Println("c=",c)


	//16进制 满16进1，已0x或0X开头
	var k int = 0x11   //=17
	fmt.Println("k=",k)



	 //二进制转换成十进制
	 //规则：从最低位开始（右边的），将每个位上的数提取出来，乘以2的（位数-1）次方，然后求和
	 //（注释第一位的位数是1  然后1-1就等于0，任何数的0次方都是1）
	 //1011  转成10进制  1*1  +  1*2  + 0*2*2  + 1*2*2*2 = 11


	//八进制转换成十进制
	 //0123  转成10进制  3*1  +  2*8  + 1*8*8  + 0*8*8*8 = 3 + 16 + 64 +0 = 83



	 //16进制转十进制
	 //0x34A =  10*1 + 4*16 + 3*16*16 = 842







	 //-------------------------------------
	 //十进制转换成二进制
	 //规则:将该数字不断除以2，直到商为0开始，然后将每步得到的余数倒过来，就是对应的二进制
	 //56转换成2进制
	 //56/2 = 28 余0
	 //28/2 = 14 余0
	 //14/2 = 7余0
	 //7/2 = 3余1
	 //3/2 =1余1
	 //最后= 111000


	 //十进制转换成八进制
	//规则:将该数字不断除以8，直到商为0开始，然后将每步得到的余数倒过来，就是对应的八进制
	// 156转换成8进制
	// 156/8 = 19 余 4
	//  19/8 = 2 余 3
	//最后= 234


	//十进制转换成十六进制
	//规则:将该数字不断除以16，直到商为0开始，然后将每步得到的余数倒过来，就是对应的十六进制
	// 356转换成16进制
	// 356 / 16 = 22 余 4
	//  22 / 16 =  1 余 6
	// 最后= 164

}
