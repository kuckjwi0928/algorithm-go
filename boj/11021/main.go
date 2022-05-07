package main

import (
	"fmt"
)

func main() {
	var T int
	_, _ = fmt.Scan(&T)

	for i := 1; i <= T; i++ {
		var A, B int
		_, _ = fmt.Scan(&A, &B)
		fmt.Printf("Case #%d: %d\n", i, A+B)
	}
}
