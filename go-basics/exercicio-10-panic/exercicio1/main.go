package main

import (
	"fmt"
	"os"
)

func openFile() {
	defer func() {
		fmt.Printf("Execução concluída")
	}()

	_, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
}

func main() {
	openFile()
}
