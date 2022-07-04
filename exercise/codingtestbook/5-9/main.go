package main

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	elements []T
	size     int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elements: make([]T, 0),
		size:     0,
	}
}

func (q *Queue[T]) Offer(element T) {
	q.elements = append(q.elements, element)
	q.size++
}

func (q *Queue[T]) Pop() (T, error) {
	if q.size == 0 {
		var noop T
		return noop, errors.New("queue is empty")
	}
	first := q.elements[0]
	q.elements = q.elements[1:]
	q.size--
	return first, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func main() {
	visited := make([]bool, 9)

	graph := [][]int{
		{},
		{2, 3, 8},
		{1, 7},
		{1, 4, 5},
		{3, 5},
		{3, 4},
		{7},
		{2, 6, 8},
		{1, 7},
	}

	bfs(graph, 1, visited)
}

func bfs(graph [][]int, v int, visited []bool) {
	visited[v] = true
	queue := NewQueue[int]()
	queue.Offer(v)
	for !queue.IsEmpty() {
		el, _ := queue.Pop()
		fmt.Println(el)
		for _, i := range graph[el] {
			if !visited[i] {
				queue.Offer(i)
				visited[i] = true
			}
		}
	}
}
