package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	_, _ = fmt.Scan(&a, &b, &c)
	arr := []int{a, b, c}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	if arr[0] == arr[1] && arr[1] == arr[2] {
		fmt.Println(10000 + arr[0]*1000)
	} else if arr[0] == arr[1] {
		fmt.Println(1000 + arr[0]*100)
	} else if arr[1] == arr[2] {
		fmt.Println(1000 + arr[1]*100)
	} else {
		fmt.Println(arr[0] * 100)
	}
}
