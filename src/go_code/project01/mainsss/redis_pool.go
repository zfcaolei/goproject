package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//创建redis连接池


var pool *redis.Pool

func init()  {
	pool = &redis.Pool{
		MaxIdle: 8 ,//最大空闲链接数
		MaxActive: 0,// 表示和数据库的最大链接数，0表示没限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","127.0.0.1:6679")
		},
	}

}


func main()  {

	v1conn := pool.Get()
	defer v1conn.Close()
	data,_ := v1conn.Do("set","namego11111","go")
	fmt.Println(data)
}