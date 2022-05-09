package main

import "fmt"

func main() {
	fmt.Println(solution([][]int{
		{1, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
	}, 1, 1, 0))
}

func solution(m [][]int, positions ...int) int {
	possiblePositions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	x, y := positions[0], positions[1]
	direction := positions[2]

	count := 0
	turnCount := 0
	visited := [50][50]int{}

	for true {
		if turnCount == 4 {
			nx := x - possiblePositions[direction][0]
			ny := y - possiblePositions[direction][1]
			if m[nx][ny] == 0 {
				x = nx
				y = ny
				turnCount = 0
			} else {
				break
			}
		}
		direction = turn(direction)
		nx := x + possiblePositions[direction][0]
		ny := y + possiblePositions[direction][1]
		if visited[nx][ny] == 0 && m[nx][ny] == 0 {
			x = nx
			y = ny
			visited[nx][ny] = 1
			turnCount = 0
			count++
		} else {
			turnCount++
		}
	}

	return count
}

func turn(direction int) int {
	direction--
	// west
	if direction == -1 {
		return 3
	}
	return direction
}
