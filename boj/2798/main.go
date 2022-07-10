package main

import "fmt"

var max = 0

func main() {
	var N, M int
	_, _ = fmt.Scan(&N, &M)

	arr := make([]int, 0)

	for i := 0; i < N; i++ {
		var num int
		_, _ = fmt.Scan(&num)
		arr = append(arr, num)
	}

	visited := make([]bool, len(arr))

	perm(visited, arr, make([]int, 3), 0, M)

	fmt.Println(max)
}

func perm(visited []bool, arr []int, check []int, index int, M int) {
	if index == 3 {
		sum := 0
		for i := 0; i < index; i++ {
			sum += check[i]
		}
		if sum > max && sum <= M {
			max = sum
		}
		return
	}

	for i := 0; i < len(arr); i++ {
		if visited[i] == true {
			continue
		}
		visited[i] = true
		check[index] = arr[i]
		perm(visited, arr, check, index+1, M)
		visited[i] = false
	}
}
