package main

import "fmt"

// https://codeforces.com/problemset/problem/1862/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		_, _ = fmt.Scan(&n, &m)
		arr := make([]string, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&arr[i])
		}

		name := "vika"
		index := 0
		found := false
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if arr[j][i] == name[index] {
					index++
					break
				}
			}
			if index == len(name) {
				found = true
				break
			}
		}

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
