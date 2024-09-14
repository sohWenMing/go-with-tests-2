package dictionary

import (
	"errors"
	"strings"
)

type Dictionary map[string]string
type DictionaryErr string

const (
	replaceValue  = "<value>"
	errNotFound   = DictionaryErr("could not find the word you were looking for : <value>")
	errWordExists = DictionaryErr("cannot add the word : <value> because it already exists in the dictionary")
)

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", generateError(errNotFound, key)
	}
	return v, nil
}

func (d Dictionary) Add(key, value string) (err error) {
	_, errFromSearch := d.Search(key)
	if errFromSearch == nil {
		return generateError(errWordExists, key)
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) (err error) {
	_, errorFromSearch := d.Search(key)
	if errorFromSearch != nil {
		return generateError(errNotFound, key)
	}
	d[key] = value
	return nil
}

func generateError(err DictionaryErr, key string) error {
	errorString := string(err)
	return errors.New(replaceErrorValue(errorString, key))
}

func replaceErrorValue(errString, key string) string {
	return strings.Replace(errString, replaceValue, key, -1)
}
