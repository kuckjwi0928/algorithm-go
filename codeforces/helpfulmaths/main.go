package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://codeforces.com/problemset/problem/339/A
func main() {
	var input string
	_, _ = fmt.Scan(&input)

	inputs := strings.Split(input, "+")

	nums := make([]int, len(inputs))

	for index, in := range inputs {
		num, _ := strconv.Atoi(in)
		nums[index] = num
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	results := make([]string, 0)

	for _, n := range nums {
		results = append(results, strconv.Itoa(n))
	}

	fmt.Println(strings.Join(results, "+"))
}
