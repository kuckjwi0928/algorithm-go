package main

import (
	"fmt"
	"strconv"
)

// https://codeforces.com/problemset/problem/1702/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		result, _ := strconv.Atoi(string((s[0]-'0'-1)+'0') + s[1:])
		fmt.Println(result)
	}
}
