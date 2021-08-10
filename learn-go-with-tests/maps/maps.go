package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrorNotFound = DictionaryErr("error en el diccionario")
	ErrWordExists = DictionaryErr("la palabra que desea ingresar ya existe")
)

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

func (e DictionaryErr) Error() string {
	return string(e)
}
