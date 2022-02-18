package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{3, 8, 9, 7, 6}, 3))
}

func Solution(A []int, K int) []int {
	length := len(A)
	for i := 0; i < K; i++ {
		last := A[length-1]
		for j := length - 1; j >= 0; j-- {
			if j == 0 {
				A[0] = last
				break
			}
			A[j] = A[j-1]
		}
	}
	return A
}
