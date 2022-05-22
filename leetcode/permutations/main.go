package main

import (
	"fmt"
)

// https://leetcode.com/problems/permutations
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

type Perm struct {
	permutations [][]int
}

func (p *Perm) Add(perm []int) {
	newPerm := make([]int, len(perm))
	copy(newPerm, perm)
	p.permutations = append(p.permutations, newPerm)
}

func permute(nums []int) [][]int {
	perm := &Perm{
		permutations: [][]int{},
	}
	permutation(nums, make([]bool, len(nums)), make([]int, len(nums)), perm, 0)
	return perm.permutations
}

func permutation(nums []int, visited []bool, arr []int, perm *Perm, index int) {
	if index == len(nums) {
		perm.Add(arr)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		arr[index] = nums[i]
		permutation(nums, visited, arr, perm, index+1)
		visited[i] = false
	}
}
