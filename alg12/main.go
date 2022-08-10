package main

import "fmt"

// Матрица Google, стр 114
// Как, и в предыдущем алгоритме, есть те же проблемы,
// плюс дополнительно, у меня не получилось получить тот же набор рейтингов, что и в книге
// Дополнительно проверил в Excel, не получается результат как в книге.
// на стр 116, в исходной матрице и в преобразованной матрице есть отличие в 4 строке,

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

func prepareMatrix(m *[][]int) (*[][]float32, *[]float32) {

	m = matrixTransposition(m)

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

func doGMatrix(f *[][]float32, d float32) {
	x, y := matrixSizeF(f)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			(*f)[i][j] = (*f)[i][j]*d + (1-d)/float32(x)
		}
	}
}

func prepareGMatrix(m *[][]int, d float32) (*[][]float32, *[]float32) {
	g, v := prepareMatrix(m)
	doGMatrix(g, d)

	return g, v

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

func printMF(m *[][]float32) {
	for _, v := range *m {
		fmt.Println(v)
	}
}

func main() {
	dampingFactor := float32(0.85)
	source := [][]int{{0, 1, 0, 0, 1, 0}, {0, 0, 1, 0, 1, 0}, {0, 1, 0, 0, 0, 1}, {1, 0, 0, 0, 1, 0}, {0, 1, 1, 0, 0, 1}, {0, 0, 0, 0, 1, 0}}

	sourceF, vector := prepareGMatrix(&source, dampingFactor)

	fmt.Println("Матрица Google")
	printMF(sourceF)

	fmt.Printf("Исходный вектор %v\n", *vector)

	steps := 1

	for i := 0; i < steps; i++ {
		vector = matrixMulti(sourceF, vector)
	}

	fmt.Printf("Итоговый рейтинг (Глубина %d): %v\n", steps, *vector)

}
