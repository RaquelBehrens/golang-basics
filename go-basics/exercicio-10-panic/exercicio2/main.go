package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile() {
	defer func() {
		fmt.Println("Execução concluída")
	}()

	file, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Conteúdo do arquivo:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		panic("Não foi possível ler o arquivo adequadamente")
	}
}

func main() {
	openFile()
}
