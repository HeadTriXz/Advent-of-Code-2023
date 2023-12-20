package utils

func Sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum
}

func GDC(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GDC(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
