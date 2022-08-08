package main

import (
	"fmt"
	"math/rand"
)

// Быстрая сортировка, стр 85

func createData(max int) *[]int {
	s := make([]int, max, max)
	for i := 0; i < max; i++ {
		s[i] = rand.Intn(1000000)
	}
	return &s
}

func sort(s *[]int, x int) {
	if len(*s) <= 1 {
		return
	}
	low := make([]int, 0, 0)
	more := make([]int, 0, 0)
	for i := 0; i < len(*s); i++ {
		if x != i {
			if (*s)[i] >= (*s)[x] {
				more = append(more, (*s)[i])
			} else {
				low = append(low, (*s)[i])
			}
		}

	}
	result := make([]int, 0, 0)
	sort(&more, len(more)/2)
	sort(&low, len(low)/2)

	result = append(result, low...)
	result = append(result, (*s)[x])
	result = append(result, more...)
	copy(*s, result)
}

func main() {
	size := 15
	s := createData(size)
	fmt.Println("Before sorting", *s)
	sort(s, len(*s)/2)
	fmt.Println("After sorting", *s)
}
