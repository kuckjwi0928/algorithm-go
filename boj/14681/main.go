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
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	if x >= 1 && y >= 1 {
		fmt.Println(1)
	} else if x < 1 && y >= 1 {
		fmt.Println(2)
	} else if x < 1 && y < 1 {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}
