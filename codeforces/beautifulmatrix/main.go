package main

import (
	"fmt"
)

func main() {
	beautifulIndex := 2

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var n int
			_, _ = fmt.Scan(&n)
			if n == 1 {
				fmt.Println(Abs(i-beautifulIndex) + Abs(j-beautifulIndex))
				return
			}
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
