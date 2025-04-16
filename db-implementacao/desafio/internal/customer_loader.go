package internal

type LoaderCustomer interface {
	Load() (cs []Customer, err error)
}
