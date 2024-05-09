package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	sum := 0
	minLength := math.MaxInt32
	for right < len(nums) {
		//fmt.Println(sum)
		sum += nums[right]
		for sum >= target { //一定是for，不是if，一直大于target时，left一直加一
			minLength = min(minLength, right-left+1)
			if minLength == 1 {
				return minLength
			}
			// fmt.Println("length =", minLength)
			// fmt.Println(left, right)
			sum -= nums[left]
			left++
		}
		right++
	}

	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
