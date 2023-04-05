package main

import "fmt"

// https://codeforces.com/problemset/problem/1607/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		keyMap := make(map[rune]int)
		var keyboard string
		_, _ = fmt.Scan(&keyboard)
		for j, k := range keyboard {
			keyMap[k] = j + 1
		}
		var keyString string
		_, _ = fmt.Scan(&keyString)
		sum := 0
		for i := 0; i < len(keyString)-1; i++ {
			sum += abs(keyMap[rune(keyString[i])] - keyMap[rune(keyString[i+1])])
		}
		fmt.Println(sum)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
