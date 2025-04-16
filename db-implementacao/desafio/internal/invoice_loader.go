package internal

type LoaderInvoice interface {
	Load() (cs []Invoice, err error)
}
