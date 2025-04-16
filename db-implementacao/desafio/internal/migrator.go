package internal

type Migrator interface {
	Migrate() (err error)
}
