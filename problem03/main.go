package main

import "fmt"

func main() {
	var i, result int
	result = 1
	for i = 1; i <= 5; i++ {
		result *= i
	}
	fmt.Println(result)
}
