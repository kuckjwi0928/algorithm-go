package main

import "fmt"

func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var l int
		_, _ = fmt.Scan(&l)
		a := make([]int, l)
		for i := 0; i < l; i++ {
			_, _ = fmt.Scan(&a[i])
		}

		left := 0
		right := l - 1

		for i := 0; i < l; i++ {
			if i%2 == 0 {
				fmt.Print(a[left], " ")
				left++
			} else {
				fmt.Print(a[right], " ")
				right--
			}
		}

		fmt.Println()
		t--
	}
}
