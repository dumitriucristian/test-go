package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("Known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("Orange")
		want := "could not find the word"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		if err.Error() != want {

			t.Errorf("got %q want %q", err, want)
		}

	})

	t.Run("Add dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
