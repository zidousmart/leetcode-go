package kit

import (
	"fmt"
)

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(list *ListNode) {
	if list == nil {
		return
	}

	var num []int
	for list != nil {
		num = append(num, list.Val)
		list = list.Next
	}
	fmt.Println(num)
}

// 头插法
func HeadInsert(v int, l *ListNode) *ListNode {
	if l == nil {
		return &ListNode{
			Val:  v,
			Next: nil,
		}
	}

	head := &ListNode{
		Val:  v,
		Next: l,
	}

	return head
}

// 尾插法
func TailInsert(v int, l *ListNode) *ListNode {
	if l == nil {
		return &ListNode{
			Val:  v,
			Next: nil,
		}
	}

	item := &ListNode{
		Val:  v,
		Next: nil,
	}

	n := l
	for n.Next != nil {
		n = n.Next
	}
	n.Next = item

	return l
}

func ReverseList(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}

	var head *ListNode // 新链表头
	cur := l
	for cur != nil {
		tmp := cur

		tmp.Next = head // 反转
		head = tmp      // 当前节点为前一个节点

		cur = cur.Next
	}

	return head
}
