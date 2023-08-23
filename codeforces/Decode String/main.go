package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://codeforces.com/problemset/problem/1729/B
func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	_, _ = fmt.Fscan(reader, &t)

	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < t; i++ {
		var n int
		var s string
		_, _ = fmt.Fscan(reader, &n, &s)

		i := n - 1

		var sb strings.Builder

		for i >= 0 {
			num := 0

			if int(s[i]-'0') == 0 {
				num = int(s[i-1] - '0')
				if s[i-2] == '1' {
					num += 10
				} else if s[i-2] == '2' {
					num += 20
				}
				sb.WriteByte(alphabet[num-1])
				i -= 3
				continue
			}

			sb.WriteByte(alphabet[int(s[i]-'0')-1])

			i--
		}

		fmt.Println(reverse(sb.String()))
	}
}

func reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
