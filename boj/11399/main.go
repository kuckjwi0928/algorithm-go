package main

import (
	"fmt"
	"sort"
)

// https://www.acmicpc.net/problem/11399
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	sum := 0
	acc := 0
	for i := 0; i < len(arr); i++ {
		sum += acc + arr[i]
		acc += arr[i]
	}
	fmt.Println(sum)
}
