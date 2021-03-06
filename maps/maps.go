package maps

import (
	"errors"
)

type Dictionary map[string]string

var errNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	// the second value is a bool, which indicates if the key was found successfully.
	definition, ok := d[word]
	if !ok {
		return "", errNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
