// 计算斐波那契数列
// 编写一个程序，计算并输出斐波那契数列的前N个数字，其中N由用户输入。
package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("请输入斐波那契数列的前N个数字：")
	fmt.Scanln(&N)

	fibonacci := []int{0, 1}

	for i := 2; i < N; i++ {
		next := fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, next)
	}

	fmt.Printf("斐波那契数列的前%d个数字是：\n", N)
	for _, num := range fibonacci {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
