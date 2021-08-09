package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("error en el diccionario")

func (d Dictionary) Search(parameter string) (string, error) {
	definition, ok := d[parameter]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}
