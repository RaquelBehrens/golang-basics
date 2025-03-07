package main

import "fmt"

func main() {
	var (
		temperatura int     = 29
		umidade     float64 = 75
		pressao     int     = 1018
	)
	fmt.Printf("A temperatura atual é %d, a umidade atual é %v, e a pressao atual é %d.\n", temperatura, umidade, pressao)
}
