package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type product struct {
	Name      string
	Price     int
	Published bool
}

func main() {

	// p := product{
	// 	Name:      "MacBook Pro",
	// 	Price:     1500,
	// 	Published: true,
	// }

	// jsonData, err := json.Marshal(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(jsonData))

	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`

	var p product

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	// server
	rt := chi.NewRouter()
	// -> endpoints
	rt.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		// set code and body
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})
	// run
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}

}
