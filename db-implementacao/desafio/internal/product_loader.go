package internal

type LoaderProduct interface {
	Load() (cs []Product, err error)
}
