package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/731/A
func main() {
	var s string
	_, _ = fmt.Scan(&s)
	pointer := 0
	sum := 0
	for i := 0; i < len(s); i++ {
		nextPointer := int(s[i] - 'a')
		if nextPointer > pointer {
			sum += min(nextPointer-pointer, pointer+26-nextPointer)
		} else {
			sum += min(pointer-nextPointer, nextPointer+26-pointer)
		}
		pointer = nextPointer
	}
	fmt.Println(sum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
