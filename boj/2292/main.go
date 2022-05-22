package main

import "fmt"

// https://www.acmicpc.net/problem/2292
func main() {
	var N int
	_, _ = fmt.Scan(&N)

	count := 1
	i := 1

	for i < N {
		i += 6 * count
		count++
	}

	fmt.Println(count)
}
