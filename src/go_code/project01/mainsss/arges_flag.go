package main

import (
	"flag"
	"fmt"
)

//flag包获取指定参数

func main()  {

	var user string
	var pwd string
	var host string
	var port int


	flag.StringVar(&user,"u","","用户名,默认为空")
	flag.StringVar(&pwd,"pwd","","密码默认为空")
	flag.StringVar(&host,"h","localhost","主机名默认localhost")
	flag.IntVar(&port,"p",3306,"端口默认为空")
	flag.Parse()
	fmt.Println(user,host,pwd,port)

	//fmt.Printf("user=%v pwd=%v host=%v port=%v",user,pwd,host,pwd)
}
