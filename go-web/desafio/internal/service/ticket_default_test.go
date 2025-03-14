package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]internal.TicketAttributes, err error) {
			t = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalTickets(context.Background())

		// assert
		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

func TestServiceTicketDefault_GetTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGetTicketsByDestinationCountry = func(country string) (t map[int]internal.TicketAttributes, err error) {
			t = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
				2: {
					Name:    "Joe",
					Email:   "joedoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}
		country := "USA"

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		tickets, err := sv.GetTicketsByDestinationCountry(context.Background(), country)

		// assert
		expectedTotal := 2
		require.NoError(t, err)
		require.Equal(t, expectedTotal, len(tickets))
	})
}
