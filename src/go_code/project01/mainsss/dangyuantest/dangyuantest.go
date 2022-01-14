package dangyuantest

//单元测试
//执行所有单元测试文件 如果需要脚本执行再当前目录 执行 go test -v
// 执行单个 go test -v dangyuantest.go add_test.go
// 执行单元测试某个函数 go test -v  -test.run TestAddnum
func Addnum(num int) (int) {
	count := 0
	for i:=0;i<=num;i++ {
		 count +=i
	}
	return  count
}

func SunNum(a1 int,a2 int)(res int)  {
	 return a1+a2
}
