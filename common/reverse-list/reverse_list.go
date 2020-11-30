package main

import (
	"fmt"
	"leetcode-go/kit"
)

func main() {
	head := kit.HeadInsert(1, nil)
	head = kit.HeadInsert(2, head)
	head = kit.HeadInsert(3, head)
	head = kit.HeadInsert(4, head)
	head = kit.HeadInsert(5, head)
	kit.PrintList(head)
	newHead := kit.ReverseList(head)
	kit.PrintList(newHead)
	fmt.Println("**********")

	tail := kit.TailInsert(1, nil)
	tail = kit.TailInsert(2, tail)
	tail = kit.TailInsert(3, tail)
	tail = kit.TailInsert(4, tail)
	tail = kit.TailInsert(5, tail)
	kit.PrintList(tail)
	newTail := kit.ReverseList(tail)
	kit.PrintList(newTail)
	fmt.Println("**********")

}
