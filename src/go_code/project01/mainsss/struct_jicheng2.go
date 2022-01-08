package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a A) Sayname() {
	fmt.Println("我的名字叫", a.Name)
}

func (a A) sayage() {
	fmt.Println("我的年龄是", a.age)
}

type B struct {
	A
}

type C struct {
	Name string
	A
}

func (c C) Sayname() {
	fmt.Println("我的名字叫", c.Name)
}

func main() {

	var b B
	b.A.Name = "曹磊"
	b.A.age = 29
	b.A.Sayname()
	b.A.sayage()

	fmt.Println()

	//可以简化b.Name
	b.Name = "测试"
	b.age = 29
	b.Sayname()
	b.sayage()

	fmt.Println()

	//如果当前结构体继承结构体有同样的属性和方法，使用简化写法就是就近原则，如果要指定继承上面的属性就要执行结构体名称
	c := new(C) //也可以写成 c:= C{}
	c.Name = "c曹磊"
	c.A.Name = "CC"
	c.age = 30
	c.A.Sayname()
	c.Sayname()
	c.sayage()
}
