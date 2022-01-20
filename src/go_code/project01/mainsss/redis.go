package main

func main()  {
	//string

	// set name carl   设置
	// get name        获取
	// setex  name 10 carl  设置十秒的时间
	// mset  name1 carl name2 mark  //设置多个
	// mget  name1 name2			//获取多个
	// del  name1				   //删除



	//hash

	//hset  caolei  name caolei   			设置
	//hset  caolei  age  男		  			设置
	//hmset	caolei  name carl age 男		设置多个

	//hget  caolei name          获取
	//hgetall  caolei            获取当前hash全部内容
	//hmget	  caolei name age	 获取多个值
	//hdel   caolei			     删除
	//hgetall caolei             获取全部

	//hlen				获取hash有多少个元素

	//hexists  caolei  name   判断hash当前字段是否存在



	//list

	//lpush city beijing shanghai tianjing    设置

	//lrange city 0 -1 取出全部

	//rpop city    从右边取出 取最先进去的
	//lpop city    从左边取出 取最晚进去的

	//lpush  city hunan  从左边插入
	//rpush  city hunan  从右边尾部插入

	//llen  city 		返回长度

	//del city     		删除



	// set  无序不能重复值

	// sadd  email 455549675@qq.com 222@qq.com  设置

	//smembers email   全部取出查看

	//sismember email 123@.com 判断值是否存在

	//srem email 123.com    删除

}