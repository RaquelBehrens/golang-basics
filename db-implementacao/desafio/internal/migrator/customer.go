package migrator

import "app/internal"

func NewCustomerMigrator(ld internal.LoaderCustomer, rd internal.RepositoryCustomer) *CustomerMigrator {
	return &CustomerMigrator{
		Loader:     ld,
		Repository: rd,
	}
}

type CustomerMigrator struct {
	Loader     internal.LoaderCustomer
	Repository internal.RepositoryCustomer
}

func (m *CustomerMigrator) Migrate() (err error) {
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
