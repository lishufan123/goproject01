// 反转字符串
// 编写一个程序，接受用户输入的字符串，然后将字符串反转并输出结果。
package main

import (
	"fmt"
)

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var input string

	fmt.Print("请输入字符串: ")
	fmt.Scanln(&input)

	reversed := reverseString(input)
	fmt.Printf("反转后的字符串: %s\n", reversed)
}
