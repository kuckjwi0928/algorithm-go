package main

import "fmt"

// https://codeforces.com/problemset/problem/1367/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for t > 0 {
		var n int
		_, _ = fmt.Scan(&n)
		a := make([]int, n)
		for i := range a {
			_, _ = fmt.Scan(&a[i])
		}

		odd, even := 0, 0
		for i := 0; i < n; i++ {
			if i%2 != a[i]%2 && a[i]%2 == 1 {
				odd++
			} else if i%2 != a[i]%2 && a[i]%2 == 0 {
				even++
			}
		}

		if odd != even {
			fmt.Println(-1)
		} else {
			fmt.Println(odd)
		}
		t--
	}
}
