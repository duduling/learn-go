package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, error := d.Search(word)
	switch error {
	case nil:
		return errWordExists
	case errNotFound:
		d[word] = def
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, error := d.Search(word)
	switch error {
	case nil:
		d[word] = def
	case errCantUpdate:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
