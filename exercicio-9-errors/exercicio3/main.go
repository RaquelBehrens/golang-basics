package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 9000

	salaryError := errors.New("error: salary is less than 10000")

	var err error
	if salary < 10000 {
		err = salaryError
	}

	if errors.Is(err, salaryError) {
		fmt.Println(err)
	} else {
		fmt.Println("Salário válido:", salary)
	}
}
