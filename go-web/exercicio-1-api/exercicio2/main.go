package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	rt := chi.NewRouter()

	rt.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var people []Person
		myDecoder := json.NewDecoder(r.Body)

		for {
			if err := myDecoder.Decode(&people); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}

		var greetings []string
		for _, person := range people {
			greetings = append(greetings, "Hello "+person.FirstName+" "+person.LastName)
		}

		response := strings.Join(greetings, "/n")

		// set code and body
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(response))
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
