package main

import (
	"fmt"
)

// 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标
func twoSum(arr []int, target int) (result [][2]int) {
	arrLastIndex := len(arr) - 1
	result = make([][2]int, 100)
	for i, value := range arr {
		for j := i + 1; j <= arrLastIndex; j++ {
			if value+arr[j] == target {
				result = append(result, [2]int{i, j})
			}
		}
	}
	return
}

func main() {
	var arr []int = []int{2, 3, 10, 12, 8, 4}
	target := 12
	result := twoSum(arr, target)
	for i, value := range result {
		if value[0]+value[1] != 0 {
			fmt.Println(i, value)
		}
	}
}
