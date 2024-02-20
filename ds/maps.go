package maps

import (
	"fmt"
)

type Dictionary map[string]string

var ErrNotFound = fmt.Errorf("could not find the word you were looking for")

func Search(dictionary map[string]string, s string) (string, error) {
	result, ok := dictionary[s]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}
