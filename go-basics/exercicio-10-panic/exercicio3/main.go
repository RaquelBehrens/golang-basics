package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	ID      int
	Name    string
	Phone   int
	Address string
}

func validateClientData(newClient Client) error {
	if newClient.Name == "" {
		return errors.New("nome do cliente é vazio")
	}
	if newClient.ID == 0 {
		return errors.New("ID do cliente é vazio")
	}
	if newClient.Phone == 0 {
		return errors.New("número de telefone do cliente é vazio")
	}
	if newClient.Address == "" {
		return errors.New("endereço do cliente é vazio")
	}
	return nil
}

func clientAlreadyExists(clients []Client, newClient Client) error {
	for _, client := range clients {
		if client.Name == newClient.Name {
			return errors.New("client already exists")
		}
	}
	return nil
}

func addClient(clients []Client, newClient Client) []Client {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Several errors were detected at runtime")
		}
		fmt.Println("End of addClient func execution")
	}()

	if err := clientAlreadyExists(clients, newClient); err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}

	if err := validateClientData(newClient); err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}

	clients = append(clients, newClient)
	return clients
}

func openFile() {
	defer func() {
		fmt.Println("End of execution")
	}()

	file, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	var clients []Client
	var id int = 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		details := strings.Split(line, ",")

		phone, _ := strconv.Atoi(strings.TrimSpace(details[1]))
		clients = addClient(clients, Client{
			ID:      id,
			Name:    details[0],
			Phone:   phone,
			Address: details[2],
		})
		id = id + 1
	}

	if scanner.Err() != nil {
		panic("Não foi possível ler o arquivo adequadamente")
	}

	fmt.Println(clients)
}

func main() {
	openFile()
}
