package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Двоичный поиск по упорядоченному списку , стр. 70

func createSlice(n int) *[]int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		x := rand.Intn(10000000)
		s = append(s, x)
	}
	sort.Ints(s)

	return &s
}

func finder(s *[]int, x int) bool {
	if len(*s) == 0 {
		return false
	}
	if len(*s) == 1 {
		if (*s)[0] == x {
			return true
		} else {
			return false
		}
	}

	i := len(*s) / 2
	if (*s)[i] == x {
		return true
	}
	if (*s)[i] > x {
		newS := (*s)[:i]
		return finder(&newS, x)
	} else {
		newS := (*s)[i:]
		return finder(&newS, x)
	}
}

func main() {
	n := 20
	s := createSlice(n)

	fmt.Println(*s)
	f1 := (*s)[6]
	f2 := 4941318 + 1

	x := finder(s, f1)
	fmt.Printf("Искомое %d, Результат поиска, %t\n", f1, x)

	x = finder(s, f2)
	fmt.Printf("Искомое %d, Результат поиска, %t\n", f2, x)

}
