package main

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标
func main() {
	arr := []int{1, 3, 5, 7, 9, 10, 11, 5, 8, 9}
	fmt.Println(twoSum(arr, 10))
	fmt.Println("------")
	fmt.Println(twoSum(arr, 30))
	fmt.Println("------")
}

func twoSum(arr []int, target int) map[int]int {
	r := make(map[int]int)

	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}

	for i, v := range arr {
		num := target - v
		if j, ok := m[num]; ok {
			r[i] = j
		}
	}

	return r
}
