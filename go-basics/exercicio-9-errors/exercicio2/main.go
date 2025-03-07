package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

func NewError() error {
	return &myError{"Salary less than 10000"}
}

func main() {
	var salary int = 149000

	var err error
	if salary < 10000 {
		err = NewError()
	}

	if errors.Is(err, &myError{}) {
		fmt.Println(err)
	} else {
		fmt.Println("Salário válido:", salary)
	}

}
