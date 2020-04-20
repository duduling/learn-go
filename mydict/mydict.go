package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not found")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, error := d.Search(word)
	switch error {
	case nil:
		return errWordExists
	case errNotFound:
		d[word] = def
	}
	return nil
}
