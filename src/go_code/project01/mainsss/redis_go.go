package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {

	conn,err := redis.Dial("tcp","127.0.0.1:6679")
	if err!=nil {
		fmt.Println("redis 连接错误 err=",err)
		return
	}
	defer conn.Close() //关闭

	//redis操作string
	_,err = conn.Do("set","namego","go")
	if err !=nil {
		fmt.Println("set redis err=",err)
	}

	//获取   //返回的数据r是interface,需要装换成字符串 redis.String
	rs,err := redis.String(conn.Do("mget","namego"))
	if err !=nil {
		fmt.Println("get  redis err=",err)
	}
	//返回的数据r是interface,需要装换成字符串
	fmt.Println(rs)

	//获取多个
	rdata,_ := redis.Strings(conn.Do("mget","ke1","11"))
	fmt.Println(rdata)





	//redis操作hash

	 //设置
	_,errs:=conn.Do("hset","go_hash","name","caoleo","age",29)

	if errs !=nil {
		fmt.Println("hset  redis err=",err)
	}

	//获取
	hashdata,_ := redis.String(conn.Do("hget","go_hash","name"))
	fmt.Println(hashdata)

	//获取hash字段为int类型的 使用String也可以
	hashdataInt,_ := redis.Int(conn.Do("hget","go_hash","age"))
	fmt.Println(hashdataInt)

	//获取多个
	hashdatas,_ := redis.Strings(conn.Do("hgetall","go_hash"))
	fmt.Println(hashdatas)




	//给任何redis类型数据设置过期时间
	//conn.Do("expire","name",29)
	conn.Do("lpush","golang_list","no1:松江",30,"no2是的",28)

}
