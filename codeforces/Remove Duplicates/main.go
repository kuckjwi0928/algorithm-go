package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/978/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	m := make(map[int]int)

	for _, a := range arr {
		m[a]++
	}

	result := make([]int, 0)

	for _, a := range arr {
		if m[a] == 1 {
			result = append(result, a)
		}
		m[a]--
	}

	fmt.Println(len(result))
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]"))
}
