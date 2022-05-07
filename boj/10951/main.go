package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		arr := strings.Split(line, " ")
		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])
		fmt.Println(a + b)
	}
}
