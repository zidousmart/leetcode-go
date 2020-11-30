package main

import (
	"fmt"
)

func main() {
	s1 := "abc123adf"
	fmt.Println(longestSubStringWithoutRepeating1(s1))
	fmt.Println(longestSubStringWithoutRepeating2(s1))
	fmt.Println("-----")

	s2 := "a123a"
	fmt.Println(longestSubStringWithoutRepeating1(s2))
	fmt.Println(longestSubStringWithoutRepeating2(s2))
	fmt.Println("-----")

	s3 := "bdd1234"
	fmt.Println(longestSubStringWithoutRepeating1(s3))
	fmt.Println(longestSubStringWithoutRepeating2(s3))
	fmt.Println("-----")

	s4 := "aaa"
	fmt.Println(longestSubStringWithoutRepeating1(s4))
	fmt.Println(longestSubStringWithoutRepeating2(s4))
	fmt.Println("-----")

}

func longestSubStringWithoutRepeating1(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		m := make(map[uint8]bool)
		m[s[i]] = true
		b := false
		for j := i + 1; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				b = true
				if max < j-i {
					max = j - i
				}
				break
			} else {
				m[s[j]] = true
			}
		}
		if !b {
			if max < len(s)-i {
				max = len(s) - i
			}
		}
	}

	return max
}

func longestSubStringWithoutRepeating2(s string) int {
	m := make(map[uint8]int)
	max := 0
	i := 0
	j := 0
	for i < len(s) {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
			i++
		} else {
			if max < i-j {
				max = i - j
			}
			for m[s[i]] != m[s[j]] {
				delete(m, s[j])
				j++
			}
			i++
			j++
		}
	}

	if max < i-j {
		max = i - j
	}

	return max
}
