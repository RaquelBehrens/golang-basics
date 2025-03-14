package service

import (
	"app/internal"
	"app/internal/repository"
	"context"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalTickets(ctx context.Context) (total int, err error) {
	mapOfTickets, err := s.rp.Get(ctx)
	total = len(mapOfTickets)
	return
}

func (s *ServiceTicketDefault) GetTicketsByDestinationCountry(ctx context.Context, country string) (total map[int]internal.TicketAttributes, err error) {
	total, err = s.rp.GetTicketsByDestinationCountry(ctx, country)
	return
}
