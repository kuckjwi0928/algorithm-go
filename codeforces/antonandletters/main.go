package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/443/A
func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	m := make(map[rune]bool)
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			m[r] = true
		}
	}
	fmt.Println(len(m))
}
