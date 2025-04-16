package migrator

import "app/internal"

func NewInvoiceMigrator(ld internal.LoaderInvoice, rd internal.RepositoryInvoice) *InvoiceMigrator {
	return &InvoiceMigrator{
		Loader:     ld,
		Repository: rd,
	}
}

type InvoiceMigrator struct {
	Loader     internal.LoaderInvoice
	Repository internal.RepositoryInvoice
}

func (m *InvoiceMigrator) Migrate() (err error) {
	all, err := m.Loader.Load()
	if err != nil {
		return
	}
	for _, s := range all {
		err = m.Repository.Save(&s)
		if err != nil {
			return
		}
	}
	return
}
