package main

import "fmt"

func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		fmt.Printf("%s%s\n", s, reverse(s))
	}
}

func reverse(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
