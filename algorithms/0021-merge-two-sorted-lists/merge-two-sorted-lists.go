package main

import (
	"fmt"
	"leetcode-go/kit"
)

// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
// 示例：
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
func main() {
	arr1 := []int{1, 3, 5, 7, 9, 11, 13, 15}
	arr2 := []int{2, 4, 6, 7, 8, 10}
	fmt.Println(mergeTwoSortedSlice(arr1, arr2))

	arr3 := []int{1, 3, 5, 7, 9, 10}
	arr4 := []int{1, 4, 6, 7, 8, 10}
	fmt.Println(mergeTwoSortedSlice(arr3, arr4))
}

func mergeTwoSortedSlice(arr1, arr2 []int) []int {
	var arr []int
	l := len(arr1)
	if l > len(arr2) {
		l = len(arr2)
	}

	i := 0
	j := 0
	for i < l && j < l {
		if arr1[i] > arr2[j] {
			arr = append(arr, arr2[j])
			j++
		} else {
			arr = append(arr, arr1[i])
			i++
		}
	}

	if i < len(arr1) {
		arr = append(arr, arr1[i:]...)
	}

	if j < len(arr2) {
		arr = append(arr, arr2[j:]...)
	}

	return arr
}

func mergeTwoSortedLists(list1, list2 *kit.ListNode) *kit.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *kit.ListNode
	var arr []int
	l := len(arr1)
	if l > len(arr2) {
		l = len(arr2)
	}

	i := 0
	j := 0
	for i < l && j < l {
		if arr1[i] > arr2[j] {
			arr = append(arr, arr2[j])
			j++
		} else {
			arr = append(arr, arr1[i])
			i++
		}
	}

	if i < len(arr1) {
		arr = append(arr, arr1[i:]...)
	}

	if j < len(arr2) {
		arr = append(arr, arr2[j:]...)
	}

	return arr
}
