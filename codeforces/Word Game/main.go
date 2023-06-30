package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	score int
	words []string
}

func newPerson() *Person {
	return &Person{
		0,
		make([]string, 0),
	}
}

func (p *Person) appendWord(word string) {
	p.words = append(p.words, word)
}

func (p *Person) addScore(score int) {
	p.score += score
}

func (p *Person) getWords() []string {
	return p.words
}

// https://codeforces.com/problemset/problem/1722/C
func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		persons := make([]*Person, 3)
		words := make(map[string]int)
		for j := 0; j < 3; j++ {
			persons[j] = newPerson()
			for k := 0; k < n; k++ {
				var word string
				_, _ = fmt.Fscan(reader, &word)
				persons[j].appendWord(word)
				words[word]++
			}
		}
		for _, person := range persons {
			for _, word := range person.getWords() {
				if words[word] == 1 {
					person.addScore(3)
				} else if words[word] == 2 {
					person.addScore(1)
				}
			}
			fmt.Print(person.score, " ")
		}
		fmt.Println()
	}
}
