package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	var input = make([]int, N+2)

	for i := 1; i <= N; i++ {
		scanner.Scan()
		text := scanner.Text()
		v, _ := strconv.Atoi(text)
		input[i] = v
	}

	stack := make([]int, 0)
	stack = append(stack, 0)
	area := 0

	for i := 1; i <= N+1; i++ {
		for len(stack) > 0 && input[stack[len(stack)-1]] > input[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area = max(area, input[top]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}

	fmt.Println(area)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
