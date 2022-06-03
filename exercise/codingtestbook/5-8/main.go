package main

import "fmt"

func main() {
	visited := make([]bool, 9)

	graph := [][]int{
		{},
		{2, 3, 8},
		{1, 7},
		{1, 4, 5},
		{3, 5},
		{3, 4},
		{7},
		{2, 6, 8},
		{1, 7},
	}

	dfs(graph, 1, visited)
}

func dfs(graph [][]int, v int, visited []bool) {
	visited[v] = true
	fmt.Println(v)
	for _, i := range graph[v] {
		if !visited[i] {
			dfs(graph, i, visited)
		}
	}
}
