package internal

type LoaderSale interface {
	Load() (cs []Sale, err error)
}
