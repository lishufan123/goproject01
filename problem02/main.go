package main

import "fmt"

func main() {
	var i, sum int
	for i = 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
