package main

import (
	"fmt"
	_ "fmt"
	"math/rand"
	"time"
)

//数组应用案例

func main()  {

	var mychars  [26]byte
	for i:=0;i<26;i++{
		mychars[i] = 'A' + byte(i)
	}

	fmt.Printf("%c ",mychars[1])

	fmt.Println("--------------------------------------------")


	//获取数组最大值下标
	abc := [...]int8{23,43,122,74,89,55}
	 dd := 0
	for key,val :=range abc{
		if abc[0]< val{
			abc[0] = val
			dd = key
		}
	}

	fmt.Println(dd)



	//求出数组的和及平均值

	arrI  := [...]int{1,2,3,4,5,6,7,8}

	var ArrSum int
	var ArrAvg float64

	for key,val :=range arrI{
		ArrSum += val
		ArrAvg = float64(ArrSum) / float64(key + 1)
	}

	fmt.Println(ArrSum,ArrAvg)


	//数组反转

	var testArr [5]int
	//生成随机数种子
	rand.Seed(time.Now().UnixNano())

	for key,_ := range testArr{
		testArr[key] = rand.Intn(100)
	}
	fmt.Println(testArr)

	temp :=0
	for i:=0;i<len(testArr) /2;i++ {
		temp = testArr[len(testArr) - 1 -i]
		testArr[len(testArr) - 1 -i] = testArr[i]
		testArr[i] = temp
	}
	
	fmt.Println(testArr)
}