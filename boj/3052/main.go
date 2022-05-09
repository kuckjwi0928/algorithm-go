package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	m := make(map[int]bool, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		n, _ := strconv.Atoi(text)
		m[n%42] = true
	}

	fmt.Println(len(m))
}
