package main
//socket 服务器端
import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {

	//这里我们循环的接收客户端发送的数据
	defer conn.Close()

	for  {
		//创建一个新的切片
		buf := make([]byte,1024)

		//1.等待客户端通过conn发送信息
		//2.如果客户端没write[发送],那么协程就阻塞在这里
		//fmt.Println("服务器在等待客户端发送信息+",conn.RemoteAddr().String())

		//从conn读取信息
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出")
			return
		}
		//显示客户端发送的内容待服务器终端
		fmt.Println(string(buf[:n]))
	}
}

func main()  {
	fmt.Println("服务器开始监听。。。。")

	//1.tcp代表网络协议
	//2.0.0.0.0:8888表示本地监听8888端口 （127.0.0.1也可以但是只代表ipv4 ip,0.0.0.0代表IPV6和IPV4 ip 都可以,最好写0.0.0.0)
	listen,err :=net.Listen("tcp","0.0.0.0:8888")
	if err!=nil {
		fmt.Println("连接错误err=",err)
		return
	}

	defer listen.Close() //延时关闭listen


	for  {
		fmt.Println("等待客户端来连接............")
		conn,err := listen.Accept()
		if err!=nil {
			fmt.Println("Accept() err=",err)
		}else{
			fmt.Println("来了一位客户端用户,ip=\n",conn.RemoteAddr().String())
		}
		go process(conn)
	}


}