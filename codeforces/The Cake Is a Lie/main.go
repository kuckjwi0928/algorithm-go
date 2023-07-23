package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, m, k int
		_, _ = fmt.Fscan(reader, &n, &m, &k)

		if n*m-1 == k {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
