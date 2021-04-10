package maps

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var errNotFound = "could not find the word you were looking for"

func (d Dictionary) Search(word string) (string, error) {
	// the second value is a bool, which indicates if the key was found successfully.
	definition, ok := d[word]
	fmt.Println("definition", definition, "ok", ok)
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
