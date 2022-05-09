package main

func sum(a []int) int {
	var r int
	for _, n := range a {
		r += n
	}
	return r
}
