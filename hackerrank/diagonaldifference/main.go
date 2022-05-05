package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	length := len(arr)
	leftDiagonal := 0
	rightDiagonal := length - 1

	var sum1 int32
	var sum2 int32

	for leftDiagonal < length && rightDiagonal > -1 {
		sum1 += arr[leftDiagonal][leftDiagonal]
		sum2 += arr[leftDiagonal][rightDiagonal]
		leftDiagonal++
		rightDiagonal--
	}

	return int32(math.Abs(float64(sum1 - sum2)))
}

func main() {
	fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}))
}
