package main

import (
	"fmt"
	"math"
)

func main() {
	var result [5]int
	max := math.MinInt
	min := math.MaxInt
	for i := 0; i < len(result); i++ {
		fmt.Printf("请输入第%d个数:", i+1)
		fmt.Scanln(&result[i])

		if result[i] > max {
			max = result[i]
		}
		if result[i] < min {
			min = result[i]
		}
	}

	fmt.Printf("最大值是: %d\n", max)
	fmt.Printf("最小值是: %d\n", min)
}
