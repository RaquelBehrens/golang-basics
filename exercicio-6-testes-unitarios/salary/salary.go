package salary

import "errors"

var (
	ErrCategoryNotFound = errors.New("salary category not found")
)

func salaryA(minutosTrabalhados float64) (resultado float64) {
	resultado = 3000.0 * (minutosTrabalhados / 60)
	return resultado + (resultado * 0.5)
}

func salaryB(minutosTrabalhados float64) (resultado float64) {
	resultado = 1500.0 * (minutosTrabalhados / 60)
	return resultado + (resultado * 0.2)
}

func salaryC(minutosTrabalhados float64) (resultado float64) {
	resultado = 1000.0 * (minutosTrabalhados / 60)
	return
}

func Salary(minutosTrabalhados float64, categoria string) (resultado float64, err error) {
	switch categoria {
	case "A":
		return salaryA(minutosTrabalhados), err
	case "B":
		return salaryB(minutosTrabalhados), err
	case "C":
		return salaryC(minutosTrabalhados), err
	}
	return 0.0, ErrCategoryNotFound
}
