package internal

// Product is an struct that represents a product
type Warehouse struct {
	// ID is the unique identifier of the product
	ID int
	// Name is the name of the product
	Name string
	// Quantity is the quantity of the product
	Address string
	// CodeValue is the code value of the product
	Telephone string
	// IsPublished is the published status of the product
	Capacity int
}

type WarehouseReport struct {
	ID           int
	Name         string
	Address      string
	Telephone    string
	Capacity     int
	ProductCount int
}
