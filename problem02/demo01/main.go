// 计算阶乘： 编写一个程序，接受用户输入的整数 N，然后计算并输出 N 的阶乘
package main

import "fmt"

func main() {
	var N int
	fmt.Print("请输入整数N：")
	fmt.Scanf("%d", &N)
	var sum int = 1
	for i := 1; i <= N; i++ {
		sum *= i
	}
	fmt.Println(sum)
}
