package main

import "fmt"

// https://codeforces.com/problemset/problem/1714/C
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s int
		_, _ = fmt.Scan(&s)
		digit := 1
		result := 0
		for j := 9; j > 0; j-- {
			if j > s {
				continue
			}
			result = result + (digit * j)
			digit *= 10
			s -= j
		}
		fmt.Println(result)
	}
}
