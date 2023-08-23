package main

import "fmt"

// https://codeforces.com/problemset/problem/1709/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var x int
		_, _ = fmt.Scan(&x)
		arr := make([]int, 3)
		for j := 0; j < 3; j++ {
			_, _ = fmt.Scan(&arr[j])
		}
		cur := arr[x-1] - 1
		c := 1
		for cur != -1 {
			fmt.Println(cur)
			cur = arr[cur] - 1
			c++
		}
		if c >= 3 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
