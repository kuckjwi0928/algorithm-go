package main

import "fmt"

// https://codeforces.com/problemset/problem/1352/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var n int
		_, _ = fmt.Scan(&n)
		var ans []int
		for i := 1; n > 0; i *= 10 {
			if n%10 != 0 {
				ans = append(ans, n%10*i)
			}
			n /= 10
		}
		fmt.Println(len(ans))
		for _, v := range ans {
			fmt.Print(v, " ")
		}
		t--
	}
}
