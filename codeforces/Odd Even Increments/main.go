package main

import "fmt"

// https://codeforces.com/problemset/problem/1669/C
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)

		check := make([]bool, 4)

		for j := 0; j < n; j++ {
			var a int
			_, _ = fmt.Scan(&a)
			if j%2 == 0 {
				if a%2 == 0 {
					check[0] = true
				} else {
					check[1] = true
				}
			} else {
				if a%2 == 0 {
					check[2] = true
				} else {
					check[3] = true
				}
			}
		}

		if (check[0] && check[1]) || (check[2] && check[3]) {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
