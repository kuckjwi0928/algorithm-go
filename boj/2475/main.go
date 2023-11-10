package main

import (
	"fmt"
	"math"
)

// https://www.acmicpc.net/problem/2475
func main() {
	arr := make([]int, 5)

	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += int(math.Pow(float64(arr[i]), 2))
	}

	fmt.Println(sum % 10)
}
