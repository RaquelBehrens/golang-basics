package main

import "fmt"

func main() {
	var palavra string
	fmt.Print("Type a word: ")
	fmt.Scanf("%s", &palavra)

	fmt.Printf("The word has %d letters!\n", len(palavra))

	for i, letter := range palavra {
		fmt.Printf("Letter number %d: %c\n", i, letter)
	}
}
