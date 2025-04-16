package migrator

import "app/internal"

func NewSaleMigrator(ld internal.LoaderSale, rd internal.RepositorySale) *SaleMigrator {
	return &SaleMigrator{
		Loader:     ld,
		Repository: rd,
	}
}

type SaleMigrator struct {
	Loader     internal.LoaderSale
	Repository internal.RepositorySale
}

func (m *SaleMigrator) Migrate() (err error) {
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
