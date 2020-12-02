package main

import (
	"fmt"
)

// 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func main() {
	s1 := "a{b(c1[]23)ad}f"
	fmt.Println(validParentheses(s1))

	s2 := "a{b(c1[]23ad}f"
	fmt.Println(validParentheses(s2))

	s3 := "a}b(c1[]23)ad}f"
	fmt.Println(validParentheses(s3))

	s4 := "aaa"
	fmt.Println(validParentheses(s4))

}

func validParentheses(s string) bool {
	m := make(map[uint8]uint8)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'
	stack := make([]uint8, 0, len(s))
	top := 0
	has := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
			top++
			has = true

		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			c, ok := m[s[i]]
			if !ok {
				return false
			}
			if stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
			top--

		}
	}

	if has && top == 0 {
		return true
	}

	return false
}
