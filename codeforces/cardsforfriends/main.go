package main

import "fmt"

// https://codeforces.com/problemset/problem/1472/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var w, h, n int
		_, _ = fmt.Scan(&w, &h, &n)

		result := 1

		for w%2 == 0 {
			result *= 2
			w /= 2
		}

		for h%2 == 0 {
			result *= 2
			h /= 2
		}

		if n <= result {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}
}
