package main

import "fmt"

// https://codeforces.com/problemset/problem/1676/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var ticket string
		_, _ = fmt.Scan(&ticket)
		sum1 := 0
		sum2 := 0
		for i := 0; i < 3; i++ {
			sum1 += int(ticket[i] - '0')
		}
		for i := 3; i < 6; i++ {
			sum2 += int(ticket[i] - '0')
		}
		if sum1 == sum2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}
