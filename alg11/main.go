package main

import "fmt"

// рейтинги в веб-графе, стр 111
// В книге нет конкретного разобранного примера,
//а есть только теоретические выкладки, но применение этой теории на примере из книги не дает описанный результат в книге :(

func matrixTransposition(m *[][]int) *[][]int {
	n := make([][]int, 0, 0)

	for i, s := range *m {

		for j, _ := range s {

			if i == 0 {
				c := make([]int, 0, 0)
				n = append(n, c)
			}

			n[j] = append(n[j], (*m)[i][j])

		}

	}

	return &n

}

func printMI(m *[][]int) {
	for _, v := range *m {
		fmt.Println(v)
	}
}

func printMF(m *[][]float32) {
	for _, v := range *m {
		fmt.Println(v)
	}
}

func matrixSizeI(m *[][]int) (x, y int) {
	x = len(*m)
	y = 0
	for _, c := range *m {
		if len(c) > y {
			y = len(c)
		}
	}
	return x, y
}

func matrixSizeF(m *[][]float32) (x, y int) {
	x = len(*m)
	y = 0
	for _, c := range *m {
		if len(c) > y {
			y = len(c)
		}
	}
	return x, y
}

func prepareMatrix(m *[][]int) (*[][]float32, *[]float32) {
	x, y := matrixSizeI(m)
	mf := make([][]float32, x, x)
	for i := 0; i < x; i++ {
		fr := make([]float32, y, y)
		mf[i] = fr
	}
	for i := 0; i < y; i++ {
		count := 0
		//fr := make([]float32, y, y)
		//mf = append(mf, fr)
		for j := 0; j < x; j++ {
			if (*m)[j][i] != 0 {
				count++
			}
		}
		for j := 0; j < x; j++ {
			if (*m)[j][i] != 0 {
				mf[j][i] = float32((*m)[j][i]) / float32(count)
			}
		}
	}

	v := make([]float32, x, x)

	for i, _ := range v {
		v[i] = 1 / float32(x)
	}

	return &mf, &v
}

func matrixMulti(mf *[][]float32, v *[]float32) *[]float32 {

	x, y := matrixSizeF(mf)
	nv := make([]float32, x, x)

	for i := 0; i < x; i++ {
		s := float32(0)
		for j := 0; j < y; j++ {
			s += (*mf)[i][j] * (*v)[j]
		}
		nv[i] = s
	}
	return &nv
}

func main() {

	source := [][]int{
		{0, 1, 0, 0, 0}, {0, 0, 1, 0, 1}, {1, 0, 0, 1, 1}, {1, 0, 1, 0, 0}, {0, 1, 1, 1, 0},
	}
	fmt.Println("Исходная матрица")
	printMI(&source)
	// Пример матрицы в книге нужно сначала транспонировать, хотя в книге об этом ничего не сказано :(
	source = *matrixTransposition(&source)
	fmt.Println("Матрица после транспонирования")
	printMI(&source)

	sourceF, vector := prepareMatrix(&source)

	fmt.Println("Матрица гиперсылок")
	printMF(sourceF)

	//startVector := createVector(sourceF)

	fmt.Printf("Исходный вектор %v\n", *vector)

	steps := 100
	for i := 0; i < steps; i++ {
		vector = matrixMulti(sourceF, vector)
	}

	fmt.Printf("Итоговый рейтинг (Глубина %d): %v\n", steps, *vector)

}
