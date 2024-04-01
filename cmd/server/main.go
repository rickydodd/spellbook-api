package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type SpellBook struct {
	Id      string
	Title   string
	Authors []string
}

func main() {
	var spellBooks []SpellBook // Maintain spellbooks in memory.

	// Set port from argument, default to port 80.
	var port string
	flag.StringVar(&port, "port", "80", "port to serve over")
	flag.Parse()

	// Routes
	http.HandleFunc("GET /v1/spellbooks", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(spellBooks)
		w.Write(data)
	})

	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
