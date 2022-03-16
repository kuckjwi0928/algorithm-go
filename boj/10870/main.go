package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	dp := make([]int, 0)
	dp = append(dp, 0)
	dp = append(dp, 1)
	dp = append(dp, 1)
	for i := 2; i < n; i++ {
		dp = append(dp, dp[i-1]+dp[i])
	}
	fmt.Println(dp[n])
}
