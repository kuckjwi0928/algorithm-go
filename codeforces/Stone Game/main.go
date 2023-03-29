package main

import "fmt"

// https://codeforces.com/problemset/problem/1538/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)

		minValue := 100
		minPos := -1
		maxValue := 0
		maxPos := -1

		for j := 0; j < n; j++ {
			var a int
			_, _ = fmt.Scan(&a)
			if a < minValue {
				minValue = a
				minPos = j
			}
			if a > maxValue {
				maxValue = a
				maxPos = j
			}
		}

		minPos, maxPos = min(minPos, maxPos), max(minPos, maxPos)

		fmt.Println(min(maxPos+1, min(minPos+1+n-maxPos, n-minPos)))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
