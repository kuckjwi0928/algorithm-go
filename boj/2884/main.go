package main

import (
	"fmt"
)

const early45 = 45

func main() {
	var h, m int
	_, _ = fmt.Scan(&h, &m)
	if m < early45 {
		m += 60
		h--
	}
	if h == -1 {
		h += 24
	}
	m -= early45
	fmt.Println(h, m)
}
