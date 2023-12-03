package main

import "fmt"

// https://www.acmicpc.net/problem/11047
func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	count := 0

	for i := n - 1; i >= 0; i-- {
		count += k / arr[i]
		k %= arr[i]
	}

	fmt.Println(count)
}
