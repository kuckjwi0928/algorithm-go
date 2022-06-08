package main

import "fmt"

// https://codeforces.com/problemset/problem/158/A
func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&nums[i])
	}

	nextRoundParticipants := 0

	for i := 0; i < n; i++ {
		if nums[k-1] <= nums[i] && nums[i] > 0 {
			nextRoundParticipants++
		}
	}

	fmt.Println(nextRoundParticipants)
}
