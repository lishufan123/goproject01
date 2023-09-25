package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("请输入一个字符串: ")
	fmt.Scanln(&input)

	// 去除字符串中的空格和特殊字符，并转换为小写
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	if isPalindrome(input) {
		fmt.Println("是回文字符串")
	} else {
		fmt.Println("不是回文字符串")
	}
}

// 判断字符串是否为回文
func isPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
