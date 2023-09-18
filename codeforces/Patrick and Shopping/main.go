package main

import "fmt"

func main() {
	var d1, d2, d3 int
	_, _ = fmt.Scan(&d1, &d2, &d3)
	totalDistance1 := d1 + d2 + d3
	totalDistance2 := 2 * (d1 + d2)
	totalDistance3 := 2 * (d1 + d3)
	totalDistance4 := 2 * (d2 + d3)
	fmt.Println(min(totalDistance1, min(totalDistance2, min(totalDistance3, totalDistance4))))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
