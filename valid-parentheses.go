package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	pairmap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) != 0 {
				//fmt.Println(string(stack))
				popvalue := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if popvalue == pairmap[s[i]] {
					continue
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]"))
}
