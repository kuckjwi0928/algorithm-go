package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://codeforces.com/problemset/problem/711/A
func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &t)
	arr := make([]string, t)
	found := false
	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		seat := s
		if !found {
			seat = strings.Replace(s, "OO", "++", 1)
		}
		if seat != s {
			found = true
		}
		arr[i] = seat
	}
	if found {
		fmt.Println("YES")
		for _, v := range arr {
			fmt.Println(v)
		}
	} else {
		fmt.Println("NO")
	}
}
