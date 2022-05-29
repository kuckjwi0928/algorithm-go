package main

import "fmt"

// https://codeforces.com/problemset/problem/71/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	for n > 0 {
		var word string
		_, _ = fmt.Scan(&word)
		if len(word) <= 10 {
			fmt.Println(word)
		} else {
			l := len(word)
			s, e := word[0], word[l-1]
			fmt.Printf("%c%d%c\n", s, len(word[1:l-1]), e)
		}
		n--
	}
}
