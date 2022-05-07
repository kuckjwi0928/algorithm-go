package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Problem: 정수 N이 입력되면 00시 00분 00초부터 N시 59분 59초까지의 모든 시각 중에서 3이 하나라도 포함되는 모든 경우의 수 구하기
 */
func main() {
	fmt.Println(time(5))
}

func time(n int) int {
	hour := 0
	count := 0
	for hour = 0; hour <= n; hour++ {
		for minute := 0; minute < 60; minute++ {
			for second := 0; second < 60; second++ {
				if hasThreeWord(strconv.Itoa(hour)) || hasThreeWord(strconv.Itoa(minute)) || hasThreeWord(strconv.Itoa(second)) {
					count++
				}
			}
		}
	}
	return count
}

func hasThreeWord(s string) bool {
	return strings.Contains(s, "3")
}
