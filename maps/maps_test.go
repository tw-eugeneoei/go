package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("word exist", func(t *testing.T) {
		expected := "this is a test"
		searchWord := "test"

		assertWordExist(t, dictionary, searchWord, expected)
	})

	t.Run("word does not exist", func(t *testing.T) {
		searchWord := "code"

		assertWordDoesNotExist(t, dictionary, searchWord)
	})
}

func TestAdd(t *testing.T) {

	t.Run("should add new word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "coder"
		meaning := "needs coffee"
		dictionary.Add(word, meaning)

		assertWordExist(t, dictionary, word, meaning)
	})

	t.Run("should not add new word if it exist", func(t *testing.T) {
		word := "coder"
		meaning := "needs coffee"
		dictionary := Dictionary{word: meaning}

		assertAddExistingWord(t, dictionary, word, meaning)
	})
}

func assertWordExist(t testing.TB, dictionary Dictionary, word string, expected string) {
	t.Helper()
	result, _ := dictionary.Search(word)

	if expected != result {
		t.Errorf("expected '%q' got '%q'", expected, result)
	}
}

func assertWordDoesNotExist(t testing.TB, dictionary Dictionary, word string) {
	t.Helper()
	_, err := dictionary.Search(word)

	if err == nil {
		t.Fatal("should throw error")
	}
}

func assertAddExistingWord(t testing.TB, dictionary Dictionary, word string, meaning string) {
	err := dictionary.Add(word, meaning)

	if err != wordAlreadyExistError {
		t.Errorf("should throw error")
	}
}
