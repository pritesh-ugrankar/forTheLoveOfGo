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
	b.Copies -= 2
	return b, nil
}

books := []Book{
	{Title: "Some title", Author: "Some author"},
	{Title: "Some other title", Author: "Some other author"},
}

books[0] = Book{Title: "new title"}
books[0].Title = "Another new title"

