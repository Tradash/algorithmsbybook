package main

import (
	"fmt"
	"math/rand"
)

// Сортировка слиянием, стр. 91

func createData(max int) *[]int {
	s := make([]int, max, max)
	for i := 0; i < max; i++ {
		s[i] = rand.Intn(1000000)
	}
	return &s
}

func splitSort(s *[]int) *[]int {
	if len(*s) > 1 {
		delitel := len(*s) / 2
		p1 := (*s)[delitel:]
		p2 := (*s)[:delitel]
		a := splitSort(&p1)
		b := splitSort(&p2)
		if len(*a) == 1 && len(*b) == 0 {
			return a
		}
		if len(*a) == 0 && len(*b) == 1 {
			return b
		}

		return merge(a, b)

	}
	return s

}

func merge(s1 *[]int, s2 *[]int) *[]int {
	//fmt.Println("Merge", *s1, *s2)
	result := make([]int, 0, 0)
	i := 0
	j := 0
	for isFinished := false; !isFinished; {

		if i >= len(*s1) && j >= len(*s2) {
			isFinished = true
			continue
		}

		if i < len(*s1) && j < len(*s2) {
			if (*s1)[i] > (*s2)[j] {
				result = append(result, (*s2)[j])
				j++

			} else {
				result = append(result, (*s1)[i])
				i++
			}
			continue

		}
		if i < len(*s1) && j >= len(*s2) {
			result = append(result, (*s1)[i])
			i++
			continue
		}
		if i >= len(*s1) && j < len(*s2) {
			result = append(result, (*s2)[j])
			j++
			continue
		}

	}
	return &result
}

func main() {
	size := 15
	s := createData(size)

	fmt.Println("Before sorting", *s)
	s = splitSort(s)

	fmt.Println("After sorting", *s)

}
