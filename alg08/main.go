package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Поразрядная сортировка, стр 81

func createData(mS int, mR int) *[]int {
	s := make([]int, mS, mS)
	min := int(math.Pow10(mR - 1))
	max := int(math.Pow10(mR) - 1)
	for i := 0; i < mS; i++ {
		x := rand.Intn(max-min) + min
		s[i] = x
	}
	return &s
}

func sort(s *[]int, mR int) {

	for i := 1; i <= mR; i++ {
		divData01 := int(math.Pow10(i))
		divData02 := int(math.Pow10(i - 1))
		tmp := make(map[int][]int)
		for j := 0; j < len(*s); j++ {
			n := ((*s)[j] % divData01) / divData02
			tmp[n] = append(tmp[n], (*s)[j])
		}
		tmpS := make([]int, 0, len(*s))
		for j := 0; j < 10; j++ {
			tmpS = append(tmpS, tmp[j]...)
		}

		copy(*s, tmpS)
	}
}

func main() {

	maxRank := 6
	maxSize := 15
	s := createData(maxSize, maxRank)
	fmt.Println("Before", *s)
	sort(s, maxRank)
	fmt.Println("After", *s)
}
