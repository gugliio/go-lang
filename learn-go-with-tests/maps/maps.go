package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("error en el diccionario")
var ErrWordExists = errors.New("la palabra que desea ingresar ya existe")

func (d Dictionary) Search(parameter string) (string, error) {
	definition, ok := d[parameter]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(parameter, value string) error {
	_, err := d.Search(parameter)

	switch err {
	case ErrorNotFound:
		d[parameter] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
