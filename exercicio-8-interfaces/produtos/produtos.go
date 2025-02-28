package produtos

const (
	smallProduct  = "small"
	mediumProduct = "medium"
	largeProduct  = "large"
)

type Product interface {
	Price() float64
}

type SmallProduct struct {
	Cost float64
}

func (p SmallProduct) Price() float64 {
	return p.Cost
}

type MediumProduct struct {
	Cost float64
}

func (p MediumProduct) Price() float64 {
	price := p.Cost + p.Cost*0.03
	return price + price*0.03
}

type LargeProduct struct {
	Cost float64
}

func (p LargeProduct) Price() float64 {
	price := p.Cost + p.Cost*0.06
	return price + 2500
}

func Factory(productType string, cost float64) Product {
	switch productType {
	case smallProduct:
		return SmallProduct{cost}
	case mediumProduct:
		return MediumProduct{cost}
	case largeProduct:
		return LargeProduct{cost}
	}
	return nil
}
