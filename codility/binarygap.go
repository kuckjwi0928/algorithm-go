package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solution(1041))
}

func Solution(N int) int {
	maxCount := 0
	zeroCount := 0
	for _, c := range strconv.FormatInt(int64(N), 2) {
		if c == 48 {
			zeroCount++
		} else {
			maxCount = max(maxCount, zeroCount)
			zeroCount = 0
		}
	}
	return maxCount
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
