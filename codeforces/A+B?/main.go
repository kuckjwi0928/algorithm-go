package main

import "fmt"

// https://codeforces.com/problemset/problem/1772/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		fmt.Println(eval(s))
	}
}

type SimpleStack struct {
	elements []int
	size     int
}

func NewSimpleStack() *SimpleStack {
	return &SimpleStack{
		elements: make([]int, 0),
	}
}

func (s *SimpleStack) push(element int) {
	s.elements = append(s.elements, element)
	s.size++
}

func (s *SimpleStack) pop() int {
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	s.size--
	return element
}

func (s *SimpleStack) empty() bool {
	return s.size == 0
}

func eval(s string) int {
	stack := NewSimpleStack()
	sum := 0

	for _, w := range s {
		if w == '+' {
			sum += stack.pop()
		} else {
			stack.push(int(w - '0'))
		}
	}

	for !stack.empty() {
		sum += stack.pop()
	}

	return sum
}
