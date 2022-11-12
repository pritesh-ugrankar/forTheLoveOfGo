package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "For the love of Go",
		Author: "John Arundel",
		Copies: 100,
	}

}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "For the love of Go",
		Author: "John Arundel",
		Copies: 0,
	}

	want := 98
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatalf("Fatal Error!! %v", err)
	}

	got := result.Copies

	if want != got {
		t.Errorf("want %d books after buying 2 books from a stock of %d books,  got %d", want, b.Copies, got)
	}

}
