package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/group-anagrams/
func main() {
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
	fmt.Println(groupAnagrams([]string{"duh", "ill"}))
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	m := make(map[string][]string)

	for _, str := range strs {
		runes := []rune(str)
		sort.SliceStable(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		key := string(runes)
		val, ok := m[key]
		if !ok {
			m[key] = []string{str}
		} else {
			m[key] = append(val, str)
		}
	}

	result := make([][]string, 0)

	for _, val := range m {
		result = append(result, val)
	}

	return result
}
