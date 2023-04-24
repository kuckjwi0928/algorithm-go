package main

import "fmt"

// https://codeforces.com/problemset/problem/1631/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)

		arr1 := make([]int, n)
		arr2 := make([]int, n)

		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&arr1[j])
		}

		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&arr2[j])
		}

		maxA := 0
		maxB := 0

		for i := 0; i < n; i++ {
			a, b := func(num1, num2 int) (int, int) {
				if num1 < num2 {
					return num1, num2
				}
				return num2, num1
			}(arr1[i], arr2[i])

			maxA = max(maxA, a)
			maxB = max(maxB, b)
		}

		fmt.Println(maxA * maxB)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
