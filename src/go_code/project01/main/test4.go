package main

import "fmt"

type  Aa struct {
	Name string
	Age string
	Child []Aa
}

type Ff  []Aa


func main()  {




	var ss = []Aa{
		{Name: "1111",Age: "1111",Child: []Aa{{Name: "aaaaa",Age: "222222"}}},
		{Name: "2222",Age: "2222"},
		{Name: "3333",Age: "333"},
		{Name: "44444",Age: "4444"},
	}
	fmt.Println(ss)

//	var  Ffs  []Aa

	 Ffs := []Aa{
		 {Name: "1111",Age: "1111",Child: []Aa{{Name: "aaaaa",Age: "222222"}}},
		 {Name: "2222",Age: "2222"},
		 {Name: "3333",Age: "333"},
		 {Name: "44444",Age: "4444"},
	 }

	 var ff = Ff{
		 {Name: "1111",Age: "1111",Child: []Aa{{Name: "aaaaa",Age: "222222"}}},
		 {Name: "2222",Age: "2222"},
		 {Name: "3333",Age: "333"},
		 {Name: "44444",Age: "4444"},
	 }
	fmt.Println(Ffs)
	fmt.Println(ff)




}