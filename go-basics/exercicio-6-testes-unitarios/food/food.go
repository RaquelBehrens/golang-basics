package food

import "errors"

var (
	ErrCategoryNotFound = errors.New("category not found")
)

const (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

func foodDog(quantidade int) int {
	return 10000 * quantidade
}

func foodCat(quantidade int) int {
	return 5000 * quantidade
}

func foodHamster(quantidade int) int {
	return 250 * quantidade
}

func foodTarantula(quantidade int) int {
	return 150 * quantidade
}

func Food(categoria string) (func(quantidade int) int, error) {
	switch categoria {
	case Dog:
		return foodDog, nil
	case Cat:
		return foodCat, nil
	case Hamster:
		return foodHamster, nil
	case Tarantula:
		return foodTarantula, nil
	}
	return nil, ErrCategoryNotFound
}
