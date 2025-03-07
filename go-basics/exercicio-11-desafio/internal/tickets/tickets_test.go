package tickets_test

import (
	"fmt"
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func init() {
	fmt.Println("aqui")
	tickets.GetTicketsFromCSV("../../tickets.csv")
}

func TestGetTotalTicketsByDestination_ValidDestination(t *testing.T) {
	// arrange
	country := "Brazil"
	expected := 45

	// act
	total, err := tickets.GetTotalTicketsByDestination(country)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, total, "Get total tickets by valid destination.")

}

func TestGetTotalTicketsByDestination_InvalidDestination(t *testing.T) {
	// arrange
	country := "non-existent country"
	expected := 0

	// act
	total, err := tickets.GetTotalTicketsByDestination(country)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, total, "Get total tickets by invalid destination.")

}

func TestGetCountByPeriod_ValidPeriod(t *testing.T) {
	// arrange
	expected := 304

	// act
	total, err := tickets.GetCountByPeriod(tickets.Night)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, total, "Get tickets count by valid period.")

}

func TestGetCountByPeriod_InvalidPeriod(t *testing.T) {
	// arrange
	expected := 0

	// act
	total, err := tickets.GetCountByPeriod("invalid period")

	// assert
	require.Error(t, err)
	require.ErrorIs(t, err, tickets.ErrInvalidPeriod)
	require.EqualError(t, err, tickets.ErrInvalidPeriod.Error())
	require.Equal(t, expected, total, "Get tickets count by invalid period.")

}

func TestAverageDestination_ValidDestination(t *testing.T) {
	// arrange
	country := "Brazil"
	total := 100
	expected := 45

	// act
	total, err := tickets.AverageDestination(country, total)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, total, "Get average tickets by valid destination.")
}

func TestAverageDestination_InvalidDestination(t *testing.T) {
	// arrange
	country := "Non-existent country"
	total := 100
	expected := 0

	// act
	total, err := tickets.AverageDestination(country, total)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, total, "Get average tickets by invalid destination.")
}
