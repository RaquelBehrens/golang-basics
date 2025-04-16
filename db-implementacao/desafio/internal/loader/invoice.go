package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

func NewInvoicesJSON(file *os.File) *InvoicesJSON {
	return &InvoicesJSON{file: file}
}

type InvoicesJSON struct {
	file *os.File
}

type InvoiceJSON struct {
	ID         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	CustomerId int     `json:"customer_id"`
	Total      float64 `json:"total"`
}

func (cj *InvoicesJSON) Load() (cs []internal.Invoice, err error) {
	var invoices []InvoiceJSON
	err = json.NewDecoder(cj.file).Decode(&invoices)
	if err != nil {
		return
	}

	for _, c := range invoices {
		cs = append(cs, internal.Invoice{
			Id: c.ID,
			InvoiceAttributes: internal.InvoiceAttributes{
				Datetime:   c.Datetime,
				CustomerId: c.CustomerId,
				Total:      c.Total,
			},
		})
	}
	return
}
