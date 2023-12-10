package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/license-key-formatting/description/
func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
}

func licenseKeyFormatting(s string, k int) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			if sb.Len()%(k+1) == k {
				sb.WriteByte('-')
			}
			sb.WriteByte(upper(s[i]))
		}
	}
	return reverse(sb.String())
}

func upper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}
	return b
}

func reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
