package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1811/A
func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int

	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var (
			n      int
			d      int
			digits string
		)
		_, _ = fmt.Fscan(reader, &n, &d, &digits)
		result := make([]byte, n+1)
		index := 0
		for _, digit := range digits {
			if int(digit-'0') < d {
				result[index] = byte(d + '0')
				index++
				d = -1
			}
			result[index] = byte(digit)
			index++
		}
		if d != -1 {
			result[index] = byte(d + '0')
		}
		fmt.Println(string(result))
	}
}
