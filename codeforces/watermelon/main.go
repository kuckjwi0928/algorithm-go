package main

import "fmt"

func main() {
	var w int
	_, _ = fmt.Scan(&w)
	if w%2 == 0 && w > 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
