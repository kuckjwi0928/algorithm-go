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
	year, _ := strconv.Atoi(scanner.Text())
	fmt.Println(year - 543)
}
