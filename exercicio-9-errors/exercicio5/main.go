package main

import (
	"fmt"
)

func salarioMensal(horasTrabalhadas int, valorHora float64) (salarioTotal float64, err error) {
	if horasTrabalhadas < 80 {
		return 0.0, fmt.Errorf("error: the worker cannot have worked less than 80 hours per month")
	}

	salarioTotal = valorHora * float64(horasTrabalhadas)
	if salarioTotal >= 150000 {
		salarioTotal += salarioTotal * 0.1
	}
	return salarioTotal, nil
}

func main() {
	salary, salaryError := salarioMensal(80, 100.0)

	if salaryError != nil {
		fmt.Println(salaryError)
	} else {
		fmt.Println("Salário válido:", salary)
	}
}
