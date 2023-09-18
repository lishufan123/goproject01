// 判断质数4
// 编写一个程序，接受用户输入的整数，然后判断该整数是否为质数，并输出结果。
package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("请输入一个整数: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("输入无效，应输入一个整数。")
	} else {
		fmt.Printf("您输入的整数是: %d\n", num)
	}
	if isPrime(num) {
		fmt.Printf("%d 是质数\n", num)
	} else {
		fmt.Printf("%d 不是质数\n", num)
	}
}

// 判断一个整数是否为质数的函数
func isPrime(num int) bool {
	if num <= 1 {
		return false // 小于等于1的数不是质数
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false // 可以被整除，不是质数
		}
	}

	return true // 不能被整除，是质数
}
