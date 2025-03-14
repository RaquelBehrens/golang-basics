package repository

import (
	"app/internal"
	"context"
)

type RepositoryTicket interface {
	Get(ctx context.Context) (map[int]internal.TicketAttributes, error)
	GetTicketsByDestinationCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error)
}
