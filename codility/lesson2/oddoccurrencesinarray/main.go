package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Solution([]int{9, 3, 9, 3, 9, 7, 9}))
}

func Solution(A []int) int {
	length := len(A)

	sort.Ints(A)

	for i := 0; i < length; i = i + 2 {
		if length == 1 || i+1 == length || A[i] != A[i+1] {
			return A[i]
		}
	}

	return -1
}
