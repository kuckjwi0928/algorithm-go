package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://www.acmicpc.net/problem/10757
func main() {
	var a, b string
	_, _ = fmt.Scan(&a, &b)

	l1 := len(a)
	l2 := len(b)

	limit := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}(l1, l2)

	if l1 < limit {
		for i := 0; i < limit-l1; i++ {
			a = "0" + a
		}
	} else if l2 < limit {
		for i := 0; i < limit-l2; i++ {
			b = "0" + b
		}
	}

	rune1 := []rune(a)
	rune2 := []rune(b)

	carry := 0
	result := make([]string, limit+1)
	for limit > 0 {
		index := limit - 1
		sum := int(rune1[index]-'0') + int(rune2[index]-'0') + carry
		result[index+1] = strconv.Itoa(sum % 10)
		carry = sum / 10
		limit--
	}

	if carry > 0 {
		result[0] = strconv.Itoa(carry)
	}

	fmt.Println(strings.Join(result, ""))
}
