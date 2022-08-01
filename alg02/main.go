package main

// Алгоритм Евклида - Метод поиска наибольшего общего делителя для двух чисел

import "fmt"

func find(x, y int) int {
	m := x % y
	fmt.Println("Шаг:", x, x/y, y, m)
	if m == 0 {
		return y
	}
	return find(y, m)
}

func main() {
	x := 136
	y := 56

	fmt.Printf("X:=%d, Y:=%d, Наименьший делитель: %d", x, y, find(x, y))

}
