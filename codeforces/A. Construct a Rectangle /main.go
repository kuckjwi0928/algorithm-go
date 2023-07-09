package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1622/A
func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int

	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var l1, l2, l3 int
		_, _ = fmt.Fscan(reader, &l1, &l2, &l3)

		if l1 == l2+l3 || l2 == l1+l3 || l3 == l1+l2 {
			fmt.Println("YES")
		} else if (l1%2 == 0 && l2 == l3) || (l2%2 == 0 && l1 == l3) || (l3%2 == 0 && l1 == l2) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
