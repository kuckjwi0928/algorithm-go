package main

import "fmt"

func main() {
	var T int

	_, _ = fmt.Scan(&T)

	for T > 0 {
		var k, n int
		_, _ = fmt.Scan(&k)
		_, _ = fmt.Scan(&n)

		arr := [15][14]int{}

		for i := 0; i < 14; i++ {
			arr[0][i] = i + 1
		}

		for i := 0; i < 15; i++ {
			arr[i][0] = 1
		}

		for i := 1; i <= k; i++ {
			for j := 1; j < n; j++ {
				arr[i][j] = arr[i-1][j] + arr[i][j-1]
			}
		}

		fmt.Println(arr[k][n-1])

		T--
	}
}
