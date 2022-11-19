package bookstore

import "errors"

type Book struct {
	Title, Author string
	Copies        int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No copies left!!")
	}
	b.Copies -= 3
	return b, nil
}
