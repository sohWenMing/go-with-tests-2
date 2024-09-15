package dictionary

import (
	"errors"
	"strings"
)

type Dictionary map[string]string
type DictionaryErr string
type Operation int

const (
	replaceValue        = "<value>"
	errNotFound         = DictionaryErr("could not find the word you were looking for : <value>")
	errNotFoundOnUpdate = DictionaryErr("could not find the word you were looking for to update: <value>")
	errWordExists       = DictionaryErr("cannot add the word : <value> because it already exists in the dictionary")
	errNotFoundOnDelete = DictionaryErr("could not find the word you were looking for the delete: <value>")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	Search Operation = iota
	Add
	Update
	Delete
)

// enum of the possible operations on the dictionary map

var operation = map[Operation]string{
	Search: "search",
	Add:    "add",
	Update: "update",
	Delete: "delete",
}

func (o Operation) String() string {
	return operation[o]
}

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", processAndGenerateError(errNotFound, key)
	}
	return v, nil
}

func (d Dictionary) Add(key, value string) (err error) {
	checkError := searchValandCheckError(d, key, Add)
	if checkError != nil {
		return checkError
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) (err error) {
	checkError := searchValandCheckError(d, key, Update)
	if checkError != nil {
		return checkError
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) (err error) {
	checkError := searchValandCheckError(d, key, Delete)
	if checkError != nil {
		return checkError
	}
	delete(d, key)
	return nil
}

func processAndGenerateError(e error, key string) error {
	if value, ok := e.(DictionaryErr); ok {
		return errors.New(replaceErrorValue(value.Error(), key))
	}
	return nil
}

func replaceErrorValue(errString, key string) string {
	return strings.Replace(errString, replaceValue, key, -1)
}

func searchValandCheckError(d Dictionary, key string, op Operation) (err error) {
	_, errFromSearch := d.Search(key)
	switch op {
	case Add:
		if errFromSearch == nil {
			return processAndGenerateError(errWordExists, key)
		}
	case Update:
		if errFromSearch != nil {
			return processAndGenerateError(errNotFoundOnUpdate, key)
		}
	case Delete:
		if errFromSearch != nil {
			return processAndGenerateError(errNotFoundOnDelete, key)
		}
	}
	return nil
}
