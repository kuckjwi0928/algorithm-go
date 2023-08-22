package main

import "fmt"

// https://codeforces.com/problemset/problem/1433/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n string
		_, _ = fmt.Scan(&n)

		l := len(n)
		count := ((int(n[0]-'0') - 1) * 10) + (l * (l + 1) / 2)
		fmt.Println(count)
	}
}
