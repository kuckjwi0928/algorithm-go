package main

import (
	"fmt"
)

// https://www.acmicpc.net/problem/2231
func main() {
	var N int
	_, _ = fmt.Scan(&N)

	digit := 0
	n := N

	for n != 0 {
		digit++
		n /= 10
	}

	start := N - digit*9

	for start <= N {
		n := start
		sum := start

		for n != 0 {
			sum += n % 10
			n /= 10
		}

		if sum == N {
			fmt.Println(start)
			return
		}

		start++
	}

	fmt.Println(0)
}
