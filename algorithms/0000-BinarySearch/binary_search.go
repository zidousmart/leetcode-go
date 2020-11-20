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
	target = 3
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
		// mid := (right + left) / 2
		mid := left + (right-left)/2
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
	right := len(arr) - 1
	for left <= right {
		// mid := (right + left) / 2
		mid := left + (right-left)/2
		if arr[mid] == target {
			right = mid - 1 // 收紧右侧边界以锁定左侧边界
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left == len(arr) || arr[left] != target {
		return -1
	}

	return left
}

// 二分查找右侧边界
// 左闭右开区间 [left, right)
func rightBoard(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		// mid := (right + left) / 2
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1 // 收紧左侧边界以锁定右侧边界
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if right < 0 || arr[right] != target {
		return -1
	}

	return right
}
