package main

// Стр 22. Алгоритмы. Самый краткий и понятный курс 2022.pdf

// Равномерное распределение элементов одного множества среди элементов другого множества

import "fmt"

func fillSlice(size int, ch string) []string {
	s := make([]string, size, size)
	for i := 0; i < size; i++ {
		s[i] = ch
	}

	return s
}

func findMod(a, b int) (int, int) {
	m := a % b
	d := (a - m) / b
	//fmt.Println("Остаток", m, "Делитель", d, a, b)
	return m, d
}

func handler(tmp [][]string, a, b int) ([][]string, int, int) {
	r, q := findMod(a, b)
	fmt.Println("a.q.b.r", a, q, b, r, tmp)

	// Оставляем справа r столбцов
	// Берем b столбцов справа добавляем слева d раз
	for i := 0; i < q; i++ {

		for j := 0; j < b; j++ {

			//fmt.Println(j, len(tmp)-i*b-b+j, i)

			tmp[j] = append(tmp[j], tmp[len(tmp)-i*b-b+j]...)
		}
	}
	// Копируем остаток
	//for i := 0; i < r; i++ {
	//	fmt.Println(i, len(tmp)-r-b+i)
	//	tmp[i] = append(tmp[i], tmp[len(tmp)-r-b+i]...)
	//}

	tmp = append(tmp[:b], tmp[b:b+r]...)

	//fmt.Println("После удаления скопированных", tmp)

	if r > 1 {
		tmp, a, b = handler(tmp, b, r)
	}

	//else {
	//	tmp = tmp[:1]
	//}

	return tmp, a, b

}

func main() {
	// а всегда больше b
	a, b := 12, 7

	sa := fillSlice(a, "X")
	sb := fillSlice(b, "*")
	//tmp := append(sa, sb...)

	tmp := make([][]string, a+b, a+b)
	// result := make([]string, 0, b+a)

	ind := 0
	for _, val := range sa {
		tmp[ind] = append(tmp[ind], val)
		ind++
	}
	for _, val := range sb {
		tmp[ind] = append(tmp[ind], val)
		ind++
	}

	fmt.Println("Исходный....", tmp, a, b)

	tmp, a, b = handler(tmp, a, b)

	fin := make([]string, 0, 0)

	for _, v := range tmp {
		fin = append(fin, v...)
	}

	fmt.Println("Финальный вариант", fin)

}
