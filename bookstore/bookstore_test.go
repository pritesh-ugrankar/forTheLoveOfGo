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
		Copies: 100,
	}

	wantToBuyCopies := 2
	wantLeftInStock := 98

	buyTheBooks, err := bookstore.Buy(b)
	if err != nil {
		t.Fatalf("Fatal Error!! %v", err)
	}

	booksBought := b.Copies - buyTheBooks.Copies
	gotLeftInStock := b.Copies - booksBought

	if wantLeftInStock != gotLeftInStock {
		t.Errorf("Want %d books left from a stock of %d books after buying %d books, but %d books were bought leaving %d books in stock.", wantLeftInStock, b.Copies, wantToBuyCopies, booksBought, gotLeftInStock)
	}

}

func TestBuyInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "For the love of Go",
		Author: "John Arundel",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Errorf("Expected non nill error, copiesBought %v", err)
	}
}
