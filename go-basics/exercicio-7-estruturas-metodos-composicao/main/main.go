package main

import (
	"fmt"
	"structs/employee"
	"structs/product"
	"time"
)

func main() {
	productFound, err := product.GetById(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(productFound.Name)
	}

	allProducts := product.GetAll()
	for _, p := range allProducts {
		fmt.Print(p.Name + ", ")
	}

	fridge := product.Product{ID: 4, Name: "Fridge", Price: 2000.0, Description: "A", Category: "Eletrônicos"}
	fridge.Save()
	fmt.Println("")

	allProducts = product.GetAll()
	for _, p := range allProducts {
		fmt.Print(p.Name + ", ")
	}

	fmt.Println()

	var newPerson = employee.Person{ID: 1, Name: "Fulano", DateOfBirth: time.Now()}
	var newEmployee = employee.Employee{ID: 10, Position: "Mascote", PersonalInformation: newPerson}
	newEmployee.PrintEmployee()
}
