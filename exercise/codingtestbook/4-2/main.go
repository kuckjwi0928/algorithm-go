package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
