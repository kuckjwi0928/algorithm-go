package main

import (
	"fmt"
)

func main() {
	var A, B int
	_, _ = fmt.Scan(&A, &B)
	var C int
	_, _ = fmt.Scan(&C)

	E := B + C
	for E >= 60 {
		if A+1 == 24 {
			A = 0
		} else {
			A++
		}
		E -= 60
	}
	fmt.Println(A, E)
}
