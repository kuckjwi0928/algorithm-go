package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Element int
	prev    *Node
	next    *Node
}

type Deque struct {
	first   *Node
	last    *Node
	maxSize int
	size    int
}

func NewDeque(maxSize int) *Deque {
	return &Deque{
		maxSize: maxSize,
	}
}

func (d *Deque) Offer(n int) {
	if d.size > d.maxSize {
		panic(errors.New("overflow queue"))
	}
	l := d.last
	newNode := &Node{Element: n}
	d.last = newNode
	if l == nil {
		d.first = newNode
	} else {
		l.next = newNode
		d.last.prev = l
	}
	d.size++
}

func (d *Deque) Poll() int {
	if d.size == 0 {
		panic(errors.New("empty queue"))
	}
	f := d.first
	d.first = f.next
	d.size--
	return f.Element
}

func (d *Deque) MoveLeft() {
	f := d.first
	l := d.last
	l.next = f
	f.prev = l
	d.first = f.next
	d.last = f
}

func (d *Deque) MoveRight() {
	f := d.first
	l := d.last
	f.prev = l
	l.next = f
	d.first = l
	d.last = l.prev
}

func (d *Deque) Index(n int) int {
	index := 0
	for node := d.first; index < d.size; node = node.next {
		if n == node.Element {
			return index
		}
		index++
	}
	return -1
}
func (d *Deque) Size() int {
	return d.size
}

func (d *Deque) String() string {
	arr := make([]int, 0)
	count := 0
	for node := d.first; count < d.size; node = node.next {
		arr = append(arr, node.Element)
		count++
	}
	return fmt.Sprintf("%v", arr)
}

// https://www.acmicpc.net/problem/1021
func main() {
	var N, M int
	_, _ = fmt.Scan(&N, &M)

	arr := make([]int, M)

	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
	}

	deque := NewDeque(N)

	for i := 1; i <= N; i++ {
		deque.Offer(i)
	}

	operatingCount := 0

	for i := 0; i < M; i++ {
		targetIndex := deque.Index(arr[i])
		if targetIndex <= deque.Size()/2 {
			for j := 0; j < targetIndex; j++ {
				deque.MoveLeft()
				operatingCount++
			}
		} else {
			for j := 0; j < deque.Size()-targetIndex; j++ {
				deque.MoveRight()
				operatingCount++
			}
		}
		deque.Poll()
	}

	fmt.Println(operatingCount)
}
