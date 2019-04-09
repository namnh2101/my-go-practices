package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	ac := []int{100, 100, 50, 40, 40, 20, 10} //len = 7
	ac2 := removeDuplicates(ac)
	//100, 50, 40, 20, 10
	pp.Println(findIndex(ac2, 40, 0, len(ac2)-1)) // len 5
}

func findIndex(ac []int, in int, l, h int) int {
	var m int
	if in < ac[h] {
		return h + 1
	}
	if in > ac[l] {
		return -1
	}

	for l < h {
		m = (l + h) / 2

		if in < ac[m] {
			l = m + 1
			continue
		}

		h = m
	}
	if ac[l] == in {
		return l
	}
	if ac[l] < in || ac[l] > in {
		return l
	}

	return -1
}

func removeDuplicates(el []int) []int {
	e := map[int]bool{}
	redis := []int{}

	for v := range el {
		if e[el[v]] == true {
		} else {
			e[el[v]] = true
			redis = append(redis, el[v])
		}
	}

	return redis
}
