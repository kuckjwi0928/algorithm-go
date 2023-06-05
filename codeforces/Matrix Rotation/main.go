package main

import "fmt"

// https://codeforces.com/problemset/problem/1772/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		matrix := make([][]int, 2)
		for j := 0; j < 2; j++ {
			matrix[j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				_, _ = fmt.Scan(&matrix[j][k])
			}
		}

		beautiful := false

		for i := 0; i < 5; i++ {
			if isBeautiful(matrix) {
				beautiful = true
				break
			}
			tmp := matrix[0][0]
			matrix[0][0] = matrix[1][0]
			matrix[1][0] = matrix[1][1]
			matrix[1][1] = matrix[0][1]
			matrix[0][1] = tmp
		}

		if beautiful {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isBeautiful(matrix [][]int) bool {
	return matrix[0][0] < matrix[0][1] && matrix[0][0] < matrix[1][0] && matrix[0][1] < matrix[1][1] && matrix[1][0] < matrix[1][1]
}
