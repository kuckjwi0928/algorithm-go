package main

import "fmt"

// https://codeforces.com/problemset/problem/1517/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)

		if n%2050 != 0 {
			fmt.Println(-1)
			continue
		}

		n /= 2050
		var ans int
		for n > 0 {
			ans += n % 10
			n /= 10
		}
		fmt.Println(ans)
	}
}
