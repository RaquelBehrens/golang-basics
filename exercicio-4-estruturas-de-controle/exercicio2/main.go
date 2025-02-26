package main

import "fmt"

func main() {
	var age int
	fmt.Print("Type your age: ")
	fmt.Scanf("%d", &age)

	if age < 22 {
		fmt.Println("Only people older than 22 years old can get a loan, sorry.")
		return
	}

	var employed string
	fmt.Print("Are you employed? Type Y or N: ")
	fmt.Scanf("%s", &employed)

	if employed == "Y" {
		var employedTime string
		fmt.Print("Are you employed for more than a year? Type Y or N: ")
		fmt.Scanf("%s", &employedTime)

		if employedTime == "Y" {
			fmt.Printf("Congratulations! You got yourself a loan :]")
		} else {
			fmt.Printf("Only employers with more than a year in the company can get a loan, sorry.")
		}

	} else {
		fmt.Println("You cannot get a loan, because you are not employed :[")
	}
}
