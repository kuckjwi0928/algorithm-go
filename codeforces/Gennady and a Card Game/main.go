package main

import "fmt"

// https://codeforces.com/problemset/problem/1097/A
func main() {
	var card string
	_, _ = fmt.Scan(&card)

	cards := make([]string, 5)
	for i := 0; i < 5; i++ {
		_, _ = fmt.Scan(&cards[i])
	}

	found := false

	for _, c := range cards {
		if c[0] == card[0] || c[1] == card[1] {
			found = true
			break
		}
	}

	if found {
		fmt.Println("YES")
	} else {

		fmt.Println("NO")
	}
}
