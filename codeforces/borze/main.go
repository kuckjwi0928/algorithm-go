package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		if s[i] == '-' && s[i+1] == '.' {
			fmt.Print(1)
			i++
		} else if s[i] == '-' && s[i+1] == '-' {
			fmt.Print(2)
			i++
		} else {
			fmt.Print(0)
		}
	}
}
