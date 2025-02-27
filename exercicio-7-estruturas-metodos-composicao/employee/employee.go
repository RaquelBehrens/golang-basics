package employee

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	ID                  int
	Position            string
	PersonalInformation Person
}

func (e Employee) PrintEmployee() {
	dob := e.PersonalInformation.DateOfBirth
	fmt.Printf("Name: %s, Date of Birth: %d/%.2d/%d, Position: %s\n", e.PersonalInformation.Name, dob.Day(), dob.Month(), dob.Year(), e.Position)
}
