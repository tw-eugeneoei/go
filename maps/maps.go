package maps

type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

const (
	wordNotFoundError     = DictionaryErr("Word not found.")
	wordAlreadyExistError = DictionaryErr("Word cannot be added, already exist.")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// * second value is a boolean which indicates if the key was found successfully
	// * meaning, ok := d[word]
	// * is ok a convention ?
	meaning, isFound := d[word]
	if isFound {
		return meaning, nil
	}
	return meaning, wordNotFoundError
}

// * An interesting property of maps is that you can modify them
// * without passing as an address to it (e.g &myMap)
func (d Dictionary) Add(word string, meaning string) error {
	_, err := d.Search(word)

	if err == wordNotFoundError {
		d[word] = meaning
	}

	if err == nil {
		return wordAlreadyExistError
	}

	return err
}

// func Search(dictionary map[string]string, word string) (string, error) {
// 	return dictionary[word], errors.New("ads")
// }
