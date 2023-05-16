package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		even := 0
		for j := 0; j < n; j++ {
			var a int
			_, _ = fmt.Fscan(reader, &a)
			if a%2 == 0 {
				even++
			}
		}
		fmt.Println(min(even, n-even))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
