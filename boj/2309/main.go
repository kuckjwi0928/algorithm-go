package main

import (
	"fmt"
	"sort"
)

// https://www.acmicpc.net/problem/2309
func main() {
	arr := make([]int, 9)
	sum := 0
	for i := 0; i < len(arr); i++ {
		var n int
		_, _ = fmt.Scan(&n)
		arr[i] = n
		sum += n
	}
	sort.Ints(arr)
	for i := 0; i < 8; i++ {
		for j := 1; j < 9; j++ {
			if i == j {
				continue
			}
			if sum-(arr[i]+arr[j]) == 100 {
				for k := 0; k < len(arr); k++ {
					if k == i || k == j {
						continue
					}
					fmt.Println(arr[k])
				}
				return
			}
		}
	}
}
