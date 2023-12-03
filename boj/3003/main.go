package main

import (
	"fmt"
	"strings"
)

// https://www.acmicpc.net/problem/3003
func main() {
	template := []int{1, 1, 2, 2, 2, 8}
	arr := make([]int, 6)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	result := make([]int, 6)
	for i := 0; i < len(template); i++ {
		result[i] = template[i] - arr[i]
	}
	fmt.Println(join(result))
}

func join(arr []int) string {
	var sb strings.Builder
	for i := 0; i < len(arr); i++ {
		if i != len(arr)-1 {
			sb.WriteString(fmt.Sprintf("%d ", arr[i]))
		} else {
			sb.WriteString(fmt.Sprintf("%d", arr[i]))
		}
	}
	return sb.String()
}
