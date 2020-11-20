package main

import (
	"fmt"
	"leetcode-go/kit"
	"strconv"
)

func main() {
	n1 := 53170
	n2 := 902
	fmt.Println(n1, "+", n2, "=", n1+n2)
	nums1 := intToSlice(n1)
	nums2 := intToSlice(n2)
	s1 := addTwoNumbers(nums1, nums2)
	fmt.Println(s1)
	fmt.Println(sliceToInt(s1))
	fmt.Println("------")
	n3 := 713
	n4 := 18938
	fmt.Println(n3, "+", n4, "=", n3+n4)
	nums3 := intToSlice(n3)
	nums4 := intToSlice(n4)
	s2 := addTwoNumbers(nums3, nums4)
	fmt.Println(s2)
	fmt.Println(sliceToInt(s2))

}

func intToSlice(n int) []int {
	s := strconv.Itoa(n)
	var r []int
	for i := 0; i < len(s); i++ {
		r = append(r, int(s[i]-'0'))
	}

	return r
}

func sliceToInt(s []int) int {
	var b []byte
	for i := 0; i < len(s); i++ {
		b = append(b, byte(s[i])+'0')
	}
	n, _ := strconv.Atoi(string(b))
	return n
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func addTwoNumbers(nums1, nums2 []int) []int {
	reverseSlice(nums1)
	reverseSlice(nums2)

	if len(nums1) > len(nums2) {
		var nums []int
		i := 0
		j := 0
		for ; i < len(nums2); i++ {
			k := nums1[i] + nums2[i] + j
			nums = append(nums, k%10)
			j = k / 10
		}
		nums = append(nums, nums1[i]+j)
		i++
		for ; i < len(nums1); i++ {
			nums = append(nums, nums1[i])
		}
		reverseSlice(nums)
		return nums
	}

	var nums []int
	i := 0
	j := 0
	for ; i < len(nums1); i++ {
		k := nums1[i] + nums2[i] + j
		nums = append(nums, k%10)
		j = k / 10
	}
	nums = append(nums, nums2[i]+j)
	i++
	for ; i < len(nums2); i++ {
		nums = append(nums, nums2[i])
	}
	reverseSlice(nums)

	return nums
}

func listAddTwoNumbers(l1 *kit.ListNode, l2 *kit.ListNode) *kit.ListNode {
	head := &ListNode{0, nil}
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = new(ListNode)
		current.Next.Val = sum % 10
		current = current.Next
	}
	return head.Next
}
