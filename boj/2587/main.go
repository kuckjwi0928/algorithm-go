package main

import (
	"fmt"
	"sort"
)

// https://www.acmicpc.net/problem/2587
func main() {
	arr := make([]int, 5)
	sum := 0

	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
		sum += arr[i]
	}

	sort.Ints(arr)

	fmt.Println(sum / len(arr))
	fmt.Println(arr[len(arr)/2])
}
