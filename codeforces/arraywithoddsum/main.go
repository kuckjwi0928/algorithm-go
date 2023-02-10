package main

import "fmt"

// https://codeforces.com/problemset/problem/1296/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&arr[j])
		}
		sum := 0
		for _, v := range arr {
			sum += v
		}
		if sum&1 == 1 {
			fmt.Println("YES")
			continue
		}
		odd := false
		even := false
		for _, v := range arr {
			if v&1 == 1 {
				odd = true
			} else {
				even = true
			}
		}
		if odd && even {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
