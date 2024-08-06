package main

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}


func (d Dictionary) Search(word string) (string,error) {
	
	if d[word] == "" {
		return "", ErrNotFound
	}
	
	return d[word], nil
}

func (d Dictionary) Add(word, definition string) (error) {
	if d[word] != "" {
		return ErrWordExists
	}
	
	d[word] = definition
	return nil
	
}