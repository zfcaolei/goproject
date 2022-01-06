package main

import "fmt"

func main()  {
	var name string
	user := make(map[string]map[string]string)
	for true {
		fmt.Println("请输入昵称")
		fmt.Scanln(&name)
		add(user,name)
		fmt.Println(user)
	}


}

func add(user map[string]map[string]string,name string)  {

	if user[name] != nil {
		user[name]["pwd"] = "8888888"
	}else{
		user[name] = make(map[string]string)
		user[name]["pwd"] = "111111"
		user[name]["nickname"] = "昵称+"+name
	}
}




