package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solution("a1"))
	fmt.Println(solution("c2"))
}

func solution(location string) int {
	arr := strings.Split(location, "")
	possiblePositions := [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}}
	mapping := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}

	x := mapping[arr[0]]
	y, _ := strconv.Atoi(arr[1])
	count := 0

	for _, pos := range possiblePositions {
		if x+pos[0] >= 1 && x+pos[0] <= 8 && y+pos[1] >= 1 && y+pos[1] <= 8 {
			count++
		}
	}

	return count
}
