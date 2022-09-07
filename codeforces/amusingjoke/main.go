package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://codeforces.com/problemset/problem/141/A
func main() {
	var guest, host, letter string

	_, _ = fmt.Scan(&guest)
	_, _ = fmt.Scan(&host)
	_, _ = fmt.Scan(&letter)

	r1 := []rune(strings.ToLower(guest + host))
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	r2 := []rune(strings.ToLower(letter))
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	if len(r1) != len(r2) {
		fmt.Println("NO")
		return
	}

	for i := 0; i < len(r2); i++ {
		if r1[i] != r2[i] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
