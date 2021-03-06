package maps

type DictionaryErr string

// * Any type with an Error() string method fulfils the error interface
// * therefore, it becomes a valid "error" type when returned in a function
func (de DictionaryErr) Error() string {
	return string(de)
}

const (
	wordNotFoundError     = DictionaryErr("Word not found.")
	wordAlreadyExistError = DictionaryErr("Word cannot be added, already exist.")
	wordCannotBeUpdated   = DictionaryErr("Cannot update word that does not exist.")
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
		return nil
	}

	if err == nil {
		return wordAlreadyExistError
	}

	return err
}

func (d Dictionary) Update(word string, meaning string) error {
	_, err := d.Search(word)

	if err == wordNotFoundError {
		return wordCannotBeUpdated
	}

	d[word] = meaning
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// func Search(dictionary map[string]string, word string) (string, error) {
// 	return dictionary[word], errors.New("ads")
// }
