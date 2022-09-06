package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/785/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	m := map[string]int{
		"Tetrahedron":  4,
		"Cube":         6,
		"Octahedron":   8,
		"Dodecahedron": 12,
		"Icosahedron":  20,
	}
	reader := bufio.NewReader(os.Stdin)

	sum := 0

	for n > 0 {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		sum += m[s]
		n--
	}

	fmt.Println(sum)
}
