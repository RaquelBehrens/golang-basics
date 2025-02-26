package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Benjamin's age: %d\n", employees["Benjamin"])

	var employeesOlderThan21 int = 0
	for _, age := range employees {
		if age > 21 {
			employeesOlderThan21 += 1
		}
	}
	fmt.Printf("%d employers are older than 21 years old.\n", employeesOlderThan21)

	employees["Frederico"] = 25
	delete(employees, "Pedro")
	for name, age := range employees {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}
