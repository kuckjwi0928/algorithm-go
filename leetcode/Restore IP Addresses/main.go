package main

import (
	"fmt"
)

// https://leetcode.com/problems/restore-ip-addresses/
func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}

func restoreIpAddresses(s string) []string {
	var result []string
	var dfs func(string, int, string)
	dfs = func(s string, n int, ip string) {
		if n == 4 {
			if len(s) == 0 {
				result = append(result, ip[:len(ip)-1])
			}
			return
		}
		for i := 1; i <= 3; i++ {
			if len(s) < i {
				break
			}
			if i > 1 && s[0] == '0' {
				break
			}
			if i == 3 && s[:3] > "255" {
				break
			}
			dfs(s[i:], n+1, ip+s[:i]+".")
		}
	}
	dfs(s, 0, "")
	return result
}
