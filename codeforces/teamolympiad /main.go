package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/490/A
func main() {
	var n int

	_, _ = fmt.Scan(&n)

	m := make(map[int][]int)

	for i := 0; i < n; i++ {
		var t int
		_, _ = fmt.Scan(&t)
		m[t] = append(m[t], i)
	}

	result := make([][]int, 0)
	count := 0

	for true {
		if len(m[1]) == 0 || len(m[2]) == 0 || len(m[3]) == 0 {
			break
		}
		result = append(result, []int{m[1][0] + 1, m[2][0] + 1, m[3][0] + 1})
		count++
		m[1] = m[1][1:]
		m[2] = m[2][1:]
		m[3] = m[3][1:]
	}

	if count == 0 {
		fmt.Print(0)
	} else {
		fmt.Println(count)
		for _, v := range result {
			fmt.Println(v[0], v[1], v[2])
		}
	}
}
