package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

func NewSalesJSON(file *os.File) *SalesJSON {
	return &SalesJSON{file: file}
}

type SalesJSON struct {
	file *os.File
}

type SaleJSON struct {
	ID        int `json:"id"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
	Quantity  int `json:"quantity"`
}

func (cj *SalesJSON) Load() (cs []internal.Sale, err error) {
	var sales []SaleJSON
	err = json.NewDecoder(cj.file).Decode(&sales)
	if err != nil {
		return
	}

	for _, c := range sales {
		cs = append(cs, internal.Sale{
			Id: c.ID,
			SaleAttributes: internal.SaleAttributes{
				ProductId: c.ProductId,
				InvoiceId: c.InvoiceId,
				Quantity:  c.Quantity,
			},
		})
	}
	return
}
