package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
	}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, error) {
	index := len(s.elements) - 1
	if index == -1 {
		var noop T
		return noop, errors.New("stack is empty")
	}
	val := s.elements[index]
	s.elements = s.elements[:index]
	return val, nil
}

func main() {
	stack := NewStack[int]()
	for i := 1; i <= 10; i++ {
		stack.Push(i)
	}
	val, _ := stack.Pop()
	fmt.Println(val)
	fmt.Println(stack)
}
