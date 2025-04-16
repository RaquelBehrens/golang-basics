package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

func NewCustomersJSON(file *os.File) *CustomersJSON {
	return &CustomersJSON{file: file}
}

type CustomersJSON struct {
	file *os.File
}

type CustomerJSON struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Condition int    `json:"condition"`
}

func (cj *CustomersJSON) Load() (cs []internal.Customer, err error) {
	var customers []CustomerJSON
	err = json.NewDecoder(cj.file).Decode(&customers)
	if err != nil {
		return
	}

	for _, c := range customers {
		cs = append(cs, internal.Customer{
			Id: c.ID,
			CustomerAttributes: internal.CustomerAttributes{
				LastName:  c.LastName,
				FirstName: c.FirstName,
				Condition: c.Condition,
			},
		})
	}
	return
}
