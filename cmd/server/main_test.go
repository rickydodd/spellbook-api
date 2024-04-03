package main

import (
	"reflect"
	"testing"
)

var spellingBook = SpellBook{
	Id:      "001",
	Title:   "Spelling Book",
	Authors: []string{"Harold Potting"},
}

var wizardsAndDoorknobs = SpellBook{
	Id:      "002",
	Title:   "Wizards and Doorknobs",
	Authors: []string{"Merlin Wizardry", "Onceler Spellcaster"},
}

var partyTricks = SpellBook{
	Id:      "003",
	Title:   "Party Tricks: A Magician's Guide to Pulling Rabbits from Hats",
	Authors: []string{"Harold Potting", "Dumbledon't Hexagonal"},
}

var spellBooks = []SpellBook{spellingBook, wizardsAndDoorknobs, partyTricks}

func TestEqual(t *testing.T) {
	testCases := []struct {
		name      string
		spellBook SpellBook
		target    SpellBook
		want      bool
	}{
		{
			name:      "spellBook, with one author, equals itself",
			spellBook: spellingBook,
			target:    spellingBook,
			want:      true,
		},
		{
			name:      "spellBook, with multiple authors, equals itself",
			spellBook: wizardsAndDoorknobs,
			target:    wizardsAndDoorknobs,
			want:      true,
		},
		{
			name:      "spellBook, with one author, does not equal other spellBook, with multiple authors",
			spellBook: spellingBook,
			target:    wizardsAndDoorknobs,
			want:      false,
		},
		{
			name:      "spellBook, with multiple authors, does not equal other spellBook, with same number multiple authors",
			spellBook: wizardsAndDoorknobs,
			target:    partyTricks,
			want:      false,
		},
	}

	for _, testCase := range testCases {
		got := testCase.spellBook.Equal(testCase.target)

		if got != testCase.want {
			t.Fatalf("%s - got: %t, want: %t", testCase.name, got, testCase.want)
		}
	}
}

func TestSpellBookById(t *testing.T) {
	testCases := []struct {
		name string
		id   string
		want SpellBook
	}{
		{
			name: "spellBook with ID exists and is found",
			id:   "001",
			want: spellingBook,
		},
		{
			name: "spellBook with ID does not exist and is not found",
			id:   "9999",
			want: (SpellBook{}),
		},
	}

	for _, testCase := range testCases {
		got := spellBookById(spellBooks, testCase.id)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Fatalf("%s - got: %s, want: %s", testCase.name, got, testCase.want)
		}
	}
}
