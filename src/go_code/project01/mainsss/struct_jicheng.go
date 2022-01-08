package main

import "fmt"

//继承

//公共共用
type Student struct {
	Name  string
	Age   int
	score float64
}

func (s *Student) SetScore(score float64) {
	s.score = score
}

func (s *Student) Show() {
	fmt.Println("学生名称:", s.Name, "学生年龄:", s.Age, "学生成绩:", s.score)
}

//小学生结构体继承学生结构体
type SmallStudent struct {
	Student
}

func (s *SmallStudent) test() {
	fmt.Println("小学生正在考试")
}

//大学生结构体继承学生结构体
type BigStudent struct {
	Student
}

func (s *BigStudent) test() {
	fmt.Println("大学生正在考试")
}

func main() {
	var smallstudent = SmallStudent{}
	smallstudent.Student.Name = "中本"
	smallstudent.Student.Age = 20
	smallstudent.test()
	smallstudent.Student.SetScore(32)
	smallstudent.Student.Show()

	var binstudent BigStudent
	binstudent.Name = "大中本"
	binstudent.Age = 32
	binstudent.test()
	binstudent.Student.SetScore(85)
	binstudent.Student.Show()
}
