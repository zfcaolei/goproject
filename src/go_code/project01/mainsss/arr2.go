package main

import "fmt"

//二维数组
func  main()  {

	//定义二维数组
	var f1 [4][6]int
	f1[1][1] = 2
	fmt.Println(f1)

	//遍历二维数组
	for _,val := range f1{
		for _,val2 := range val{
			fmt.Println(val2)
		}
		fmt.Println()
	}

	fmt.Println("---------------------------------------------------")

	//使用方式2
	var d2 = [2][2]int{{11,22},{33,44}}
	fmt.Println(d2)

	d3 := [...][2]int{{11,22},{33,44}}
	fmt.Println(d3)





	var scoreArr [2][4]float64

	for key1,val1 := range scoreArr{
		for key2,_ :=range val1{
			fmt.Println("请输入",key1+1,"班第",key2+1,"个同学成绩")
			fmt.Scanln(&scoreArr[key1][key2])
		}
	}

	fmt.Println(scoreArr)
	allcountNum :=0.0
	allcountAvg :=0
	for key,val := range scoreArr{
		countNum := 0.0
		countAvg := 0.0
		for key2,val2 := range val{
			countNum +=val2
			countAvg = countNum/float64(key2+1)
			allcountNum += val2
			allcountAvg++
		}
		fmt.Println("第",key+1,"个班平均成绩为",countAvg)
	}

	fmt.Println("总成绩为",allcountNum,"总人数为",allcountAvg,"总平均分",allcountNum/float64(allcountAvg))
}