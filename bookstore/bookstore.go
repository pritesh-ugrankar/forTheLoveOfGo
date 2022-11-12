package bookstore

type Book struct {
	Title, Author string
	Copies        int
}

func Buy(b Book) Book {
	b.Copies -= 3
	return b
}
