// 计算平均值
// 编写一个程序，接受一组数字作为输入，并计算它们的平均值，并输出结果。
package main

import "fmt"

func main() {
	var num [5]int
	var sum int

	for i := 0; i < len(num); i++ {
		fmt.Printf("请输入第%d个数：", i+1)
		_, err := fmt.Scanf("%d", &num[i])

		if err != nil {
			fmt.Println("输入无效，请输入一个整数。")
			i-- // 重新尝试读取当前位置的整数
		} else {
			fmt.Printf("您输入的整数是: %d\n", num[i])
			sum += num[i]
		}
	}

	result := float64(sum) / float64(len(num))
	fmt.Printf("平均值是: %f\n", result)
}
