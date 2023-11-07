package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{10}))
}

func maximumGap(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums)-1; i++ {
		result = max(nums[i+1]-nums[i], result)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
