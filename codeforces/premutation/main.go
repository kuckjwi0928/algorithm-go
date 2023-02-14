package main

import "fmt"

// https://codeforces.com/problemset/problem/1790/C
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		arr := make([][]int, n)
		freq := make([]int, n+1)
		for j := 0; j < n; j++ {
			arr[j] = make([]int, n-1)
			for k := 0; k < n-1; k++ {
				_, _ = fmt.Scan(&arr[j][k])
			}
			freq[arr[j][0]]++
		}

		result := make([]int, n)
		index := 0

		for j := 1; j <= n; j++ {
			if freq[j] == n-1 {
				result[index] = j
				index++
				break
			}
		}

		for i := 0; i < n; i++ {
			if arr[i][0] != result[0] {
				for j := 0; j < n-1; j++ {
					result[index] = arr[i][j]
					index++
				}
				break
			}
		}

		for _, v := range result {
			fmt.Printf("%d ", v)
		}

		fmt.Println()
	}
}
