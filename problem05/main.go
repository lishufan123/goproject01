package main

import (
	"fmt"
)

func main() {
	var n, num, max, min int

	// 获取整数的数量
	fmt.Print("请输入整数的数量: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	// 初始化最大值和最小值
	fmt.Print("请输入第1个整数: ")
	_, err = fmt.Scanln(&num)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}
	max = num
	min = num

	// 依次输入剩余的整数并更新最大值和最小值
	for i := 2; i <= n; i++ {
		fmt.Printf("请输入第%d个整数: ", i)
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("输入错误:", err)
			return
		}

		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	fmt.Printf("最大值是: %d\n", max)
	fmt.Printf("最小值是: %d\n", min)
}
