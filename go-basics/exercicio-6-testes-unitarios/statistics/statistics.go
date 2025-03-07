package statistics

import "errors"

var (
	ErrCategoryNotFound = errors.New("category not found")
)

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

// const minUint = 0
const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

// Quarto exercício
func calcMinimum(numbers ...int) int {
	var min int = maxInt
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func calcMaximum(numbers ...int) int {
	var max int = minInt
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func calcAverage(numbers ...int) int {
	var max int = 0
	for _, number := range numbers {
		max += number
	}
	return max / len(numbers)
}

func Operation(categoria string) (func(numbers ...int) int, error) {
	switch categoria {
	case Minimum:
		return calcMinimum, nil
	case Average:
		return calcAverage, nil
	case Maximum:
		return calcMaximum, nil
	}
	return nil, ErrCategoryNotFound
}
