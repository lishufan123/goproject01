package main

import "fmt"

func main() {
	var i int
	fmt.Print("请输入一个整数: ")
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println("输入无效，应输入一个整数。")
		return
	}

	if i%2 == 0 {
		fmt.Println("所输入的数是偶数。")
	} else {
		fmt.Println("所输入的数是奇数。")
	}
}
