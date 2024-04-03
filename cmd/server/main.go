package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
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

	http.HandleFunc("GET /v1/spellbook/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		spellBook := spellBookById(spellBooks, id)
		if spellBook.Equal(SpellBook{}) {
			io.WriteString(w, "SpellBook not found")
		}
		data, _ := json.Marshal(spellBook)
		w.Write(data)
	})

	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

// Equal returns whether two SpellBooks are equivalent.
// The order of authors has semantic relevance to equivlance, currently.
func (s SpellBook) Equal(target SpellBook) bool {
	if len(s.Authors) != len(target.Authors) {
		return false
	}

	if s.Id != target.Id || s.Title != target.Title {
		return false
	}

	for i := 0; i < len(s.Authors); i++ {
		if s.Authors[i] != target.Authors[i] {
			return false
		}
	}

	return true
}

// spellBookById searches a slice of SpellBooks for a specific SpellBook by
// the Id field.
func spellBookById(spellBooks []SpellBook, targetId string) SpellBook {
	var target SpellBook

	for i := 0; i < len(spellBooks); i++ {
		if spellBooks[i].Id == targetId {
			return spellBooks[i]
		}
	}

	return target
}
