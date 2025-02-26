package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	// Primeiro exercício
	var salario = 151000.0
	var resultado float64 = salarioLiquido(salario)
	fmt.Println(resultado)

	// Segundo exercício
	var media float64 = calculaMedia(10, 9, 3, 5)
	fmt.Println(media)

	// Terceiro exercício
	var minutosTrabalhados float64 = 10500.0
	fmt.Println(calculaSalario(minutosTrabalhados, "C"))

	// Quarto exercício
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)

	// Exercício 5
	animalDog, msg := animal(dog)
	if err != msg {
		fmt.Println(err)
	}
	animalCat, msg := animal(cat)
	if err != msg {
		fmt.Println(err)
	}

	var amount int
	amount += animalDog(10)
	amount += animalCat(10)
	fmt.Printf("%.1fkg de comida.\n", float64(amount)/1000)
}

// Primeiro exercício
func salarioLiquido(salario float64) (resultado float64) {
	resultado = salario
	if salario >= 50000 {
		resultado -= (salario * 0.17)
	}
	if salario >= 150000 {
		resultado -= (salario * 0.10)
	}
	return
}

// Segundo exercício
func calculaMedia(notas ...float64) float64 {
	var sum float64 = 0
	for _, nota := range notas {
		sum += nota
	}
	return sum / float64(len(notas))
}

// Terceiro exercício
func calculaSalarioA(minutosTrabalhados float64) (resultado float64) {
	resultado = 3000.0 * (minutosTrabalhados / 60)
	return resultado + (resultado * 0.5)
}

func calculaSalarioB(minutosTrabalhados float64) (resultado float64) {
	resultado = 1500.0 * (minutosTrabalhados / 60)
	return resultado + (resultado * 0.2)
}

func calculaSalarioC(minutosTrabalhados float64) (resultado float64) {
	resultado = 1000.0 * (minutosTrabalhados / 60)
	return
}

func calculaSalario(minutosTrabalhados float64, categoria string) (resultado float64) {
	switch categoria {
	case "A":
		return calculaSalarioA(minutosTrabalhados)
	case "B":
		return calculaSalarioB(minutosTrabalhados)
	case "C":
		return calculaSalarioC(minutosTrabalhados)
	}
	return 0.0
}

// const minUint = 0
const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

// Quarto exercício
func calcMinimum(numbers ...int) int {
	var min int = maxInt
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func calcMaximum(numbers ...int) int {
	var max int = minInt
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func calcAverage(numbers ...int) int {
	var max int = 0
	for _, number := range numbers {
		max += number
	}
	return max / len(numbers)
}

func operation(categoria string) (func(numbers ...int) int, error) {
	switch categoria {
	case maximum:
		return calcMinimum, nil
	case average:
		return calcAverage, nil
	case minimum:
		return calcMaximum, nil
	}
	return nil, errors.New("category not found")
}

// Exercício 5
func comidaCao(quantidade int) int {
	return 10000 * quantidade
}

func comidaGato(quantidade int) int {
	return 5000 * quantidade
}

func comidaHamster(quantidade int) int {
	return 250 * quantidade
}

func comidaTarantula(quantidade int) int {
	return 150 * quantidade
}

func animal(categoria string) (func(quantidade int) int, error) {
	switch categoria {
	case dog:
		return comidaCao, nil
	case cat:
		return comidaGato, nil
	case hamster:
		return comidaHamster, nil
	case tarantula:
		return comidaTarantula, nil
	}
	return nil, errors.New("category not found")
}
