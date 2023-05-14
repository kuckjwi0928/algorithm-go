package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1611/A
func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		if int(s[len(s)-1]-'0')%2 == 0 {
			fmt.Println(0)
			continue
		}

		if int(s[0]-'0')%2 == 0 {
			fmt.Println(1)
			continue
		}

		found := false
		for _, w := range s {
			if int(w-'0')%2 == 0 {
				found = true
				break
			}
		}

		if found {
			fmt.Println(2)
		} else {
			fmt.Println(-1)
		}
	}
}
