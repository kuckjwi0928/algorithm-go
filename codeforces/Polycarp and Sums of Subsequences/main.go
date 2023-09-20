package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		b := make([]int, 7)
		for i := 0; i < 7; i++ {
			_, _ = fmt.Fscan(reader, &b[i])
		}
		fmt.Println(b[0], b[1], b[6]-b[0]-b[1])
	}
}
