package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 7, 10, 15, 30, 32, 33, 68, 79}
	target := 10
	fmt.Println(binarySearch(arr, target))
	target = 19
	fmt.Println(binarySearch(arr, target))
	fmt.Println("------")

	arrLeft := []int{1, 3, 7, 10, 15, 15, 15, 33, 68, 79}
	target = 15
	fmt.Println(leftBoard(arrLeft, target))
	target = 19
	fmt.Println(leftBoard(arrLeft, target))
	fmt.Println("------")

	target = 15
	fmt.Println(rightBoard(arrLeft, target))
	target = 19
	fmt.Println(rightBoard(arrLeft, target))
	fmt.Println("------")
}

// 二分查找
// 左闭右闭区间 [left, right]
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (right + left) / 2
		// mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// 二分查找左侧边界
// 左闭右开区间 [left, right)
func leftBoard(arr []int, target int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := (right + left) / 2
		// mid := left + (right-left)/2
		if arr[mid] == target {
			right = mid // 收紧右侧边界以锁定左侧边界
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == len(arr) {
		return -1
	}
	if arr[left] == target {
		return left
	}

	return -1
}

// 二分查找右侧边界
// 左闭右开区间 [left, right)
func rightBoard(arr []int, target int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := (right + left) / 2
		// mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1 // 收紧左侧边界以锁定右侧边界
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == 0 {
		return -1
	}
	if arr[left-1] == target {
		return left - 1
	}

	return -1
}
