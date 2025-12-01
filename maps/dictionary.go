package main

const (
	ErrNotFound         = DictionaryErr("could not find word")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if ok {
		return definition, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word, definition string) error {
	// check if the key exists
	_, err := d.Search(word)

	switch err {
	// if does not exist, add it
	case ErrNotFound:
		d[word] = definition
	// if key exists return error
	case nil:
		return ErrWordExists
	// if some other error, propagate
	default:
		return err
	}

	// return nil if word was added
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return nil
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return nil
	}

	return nil
}
