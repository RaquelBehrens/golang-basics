package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 9000

	salaryError := fmt.Errorf("the minimum taxable amount is 150,000, and the salary entered is: %d", salary)

	var err error
	if salary < 10000 {
		err = fmt.Errorf("error: %w", salaryError)
	}

	if errors.Is(err, salaryError) {
		fmt.Println(err)
	} else {
		fmt.Println("Salário válido:", salary)
	}
}
