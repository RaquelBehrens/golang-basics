package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.GetTicketsFromCSV("tickets.csv")

	total, err := tickets.GetTotalTicketsByDestination("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)

	// total, err := tickets.GetCountByPeriod(tickets.Evening)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(total)

	// total, err := tickets.AverageDestination("Brazil", 100)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(total)
}
