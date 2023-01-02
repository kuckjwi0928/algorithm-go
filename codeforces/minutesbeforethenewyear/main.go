package main

import (
	"fmt"
	"time"
)

// https://codeforces.com/problemset/problem/1283/A

const timeFormat = "15:04"

func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var h, m int
		_, _ = fmt.Scan(&h, &m)

		parsedTime, _ := time.Parse(timeFormat, fmt.Sprintf("%02d:%02d", h, m))
		diffTargetTime, _ := time.Parse(timeFormat, "23:59")

		fmt.Println(diffTargetTime.Sub(parsedTime).Minutes() + 1)

		t--
	}
}
