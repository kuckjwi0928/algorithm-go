package main

import "fmt"

// https://codeforces.com/problemset/problem/1353/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var n, k int
		_, _ = fmt.Scan(&n, &k)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&a[i])
		}
		b := make([]int, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&b[i])
		}
		for i := 0; i < k; i++ {
			minIndex := 0
			maxIndex := 0
			for j := 0; j < n; j++ {
				if a[j] < a[minIndex] {
					minIndex = j
				}
				if b[j] > b[maxIndex] {
					maxIndex = j
				}
			}
			if a[minIndex] > b[maxIndex] {
				break
			}
			a[minIndex], b[maxIndex] = b[maxIndex], a[minIndex]
		}
		sum := 0
		for i := 0; i < n; i++ {
			sum += a[i]
		}
		fmt.Println(sum)
		t--
	}
}
