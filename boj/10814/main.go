package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Member struct {
	Age  int
	Name string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	members := make([]*Member, 0)

	for i := 0; i < N; i++ {
		scanner.Scan()
		text := scanner.Text()
		member := strings.Split(text, " ")

		age, _ := strconv.Atoi(member[0])
		name := member[1]

		members = append(members, &Member{Age: age, Name: name})
	}

	sort.SliceStable(members, func(i, j int) bool {
		return members[i].Age < members[j].Age
	})

	for _, member := range members {
		fmt.Println(member.Age, member.Name)
	}
}
