package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{1}))
	fmt.Println(jump([]int{2, 1}))
	fmt.Println(jump([]int{1, 3, 2}))
	fmt.Println(jump([]int{3, 2, 1}))
}

func jump(nums []int) int {
	length := len(nums)
	if length == 0 || length == 1 {
		return 0
	}
	curr, next, i := nums[0], nums[0], 0
	count := 1
	for i < length {
		if i > curr {
			count++
			curr = next
		}
		next = max(next, i+nums[i])
		i++
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
