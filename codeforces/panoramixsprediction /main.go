package main

import "fmt"

// https://codeforces.com/problemset/problem/80/A
func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	for i := n + 1; ; i++ {
		if isPrime(i) {
			if i == m {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
			break
		}
	}
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
