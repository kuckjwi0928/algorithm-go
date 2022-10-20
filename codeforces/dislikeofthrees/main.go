package main

import "fmt"

// https://codeforces.com/problemset/problem/1560/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	arr := make([]int, 1001)

	j := 1

	for i := 1; i <= 1666; i++ {
		if i%3 != 0 && i%10 != 3 {
			arr[j] = i
			j++
		}
	}

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		fmt.Println(arr[n])
	}
}
