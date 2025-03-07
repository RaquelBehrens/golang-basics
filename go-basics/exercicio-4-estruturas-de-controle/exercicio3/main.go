package main

import "fmt"

func main() {
	var month int
	fmt.Print("Type the month number: ")
	fmt.Scanf("%d", &month)

	var months = [12]string{"January", "Fabruary", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	fmt.Printf("%d, %s\n", month, months[month-1])
}
