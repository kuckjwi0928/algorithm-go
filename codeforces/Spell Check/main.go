package main

import "fmt"

func main() {
	var t int
	_, _ = fmt.Scan(&t)

l:
	for i := 0; i < t; i++ {
		var n int
		var s string

		_, _ = fmt.Scan(&n)
		_, _ = fmt.Scan(&s)

		if n != 5 {
			fmt.Println("NO")
			continue
		}

		checker := map[rune]bool{
			'T': false,
			'i': false,
			'm': false,
			'u': false,
			'r': false,
		}

		for _, w := range s {
			checker[w] = true
		}

		for _, v := range checker {
			if !v {
				fmt.Println("NO")
				continue l
			}
		}

		fmt.Println("YES")
	}
}
