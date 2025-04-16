package migrator

import "app/internal"

func NewProductMigrator(ld internal.LoaderProduct, rd internal.RepositoryProduct) *ProductMigrator {
	return &ProductMigrator{
		Loader:     ld,
		Repository: rd,
	}
}

type ProductMigrator struct {
	Loader     internal.LoaderProduct
	Repository internal.RepositoryProduct
}

func (m *ProductMigrator) Migrate() (err error) {
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
