package main

import "fmt"

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

func main() {
	var salary int = 149000

	if salary < 150000 {
		e := myError{"The salary entered does not reach the taxable minimum"}
		result := e.Error()
		fmt.Println(result)
	} else {
		fmt.Println("Must pay tax")
	}

}
