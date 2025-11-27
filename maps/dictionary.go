package main

import "errors"

var ErrNotFound = errors.New("could not find word")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if ok {
		return definition, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
