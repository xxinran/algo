package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	hashmap := make(map[byte]bool)
	maxlength := 0

	if len(s) <= 1 {
		return len(s)
	}

	for right < len(s) {
		if hashmap[s[right]] {

			fmt.Println(s[left])
			fmt.Println(s[right])
			delete(hashmap, s[left]) // right不行
			left++
		} else {
			hashmap[s[right]] = true
			maxlength = max(maxlength, right-left+1)
			right++
		}
	}

	return maxlength

}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
