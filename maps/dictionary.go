package main

import "errors"

type Dictionary map[string]string


var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string,error) {
	
	if d[word] == "" {
		return "", ErrNotFound
	}
	
	return d[word], nil
}