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
		err := dictionary.Add(word, meaning)

		if err != nil {
			t.Fatal("should not throw error")
		}

		assertWordExist(t, dictionary, word, meaning)
	})

	t.Run("should not add new word if it exist", func(t *testing.T) {
		word := "coder"
		meaning := "needs coffee"
		dictionary := Dictionary{word: meaning}

		err := dictionary.Add(word, meaning)

		if err != wordAlreadyExistError {
			t.Errorf("should throw error")
		}
		// assertAddExistingWord(t, dictionary, word, meaning)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("should update word with new meaning", func(t *testing.T) {
		word := "coder"
		meaning := "needs coffee"
		dictionary := Dictionary{word: meaning}
		newMeaning := "eat sleep code repeat"

		err := dictionary.Update(word, newMeaning)

		if err != nil {
			t.Fatal("should not throw error")
		}

		assertWordExist(t, dictionary, word, newMeaning)
	})

	t.Run("should throw error if word to update does not exist", func(t *testing.T) {
		word := "coder"
		meaning := "needs coffee"
		dictionary := Dictionary{}

		err := dictionary.Update(word, meaning)

		if err == nil {
			t.Errorf("should throw error")
		}
		// assertCannotUpdateWordThatDoesNotExist(t, dictionary, word, meaning)
	})
}

func TestDelete(t *testing.T) {
	word := "coder"
	meaning := "needs coffee"
	dictionary := Dictionary{word: meaning}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)

	if err != wordNotFoundError {
		t.Errorf("expect %q to be deleted", word)
	}
}

func assertWordExist(t testing.TB, dictionary Dictionary, word string, expected string) {
	t.Helper()
	result, _ := dictionary.Search(word)

	if expected != result {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func assertWordDoesNotExist(t testing.TB, dictionary Dictionary, word string) {
	t.Helper()
	_, err := dictionary.Search(word)

	if err == nil {
		t.Fatal("should throw error")
	}
}

// func assertAddExistingWord(t testing.TB, dictionary Dictionary, word string, meaning string) {
// 	err := dictionary.Add(word, meaning)

// 	if err != wordAlreadyExistError {
// 		t.Errorf("should throw error")
// 	}
// }

// func assertCannotUpdateWordThatDoesNotExist(t testing.TB, dictionary Dictionary, word string, meaning string) {
// 	err := dictionary.Update(word, meaning)

// 	if err == nil {
// 		t.Errorf("should throw error")
// 	}
// }
