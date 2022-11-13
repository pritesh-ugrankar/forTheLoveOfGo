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

	leftInStockAfterBuy := 98
	buyTheBooks, err := bookstore.Buy(b)
	booksBought := b.Copies - leftInStockAfterBuy
	if err != nil {
		t.Fatalf("Fatal Error!! %v", err)
	}

	got := buyTheBooks.Copies

	if leftInStockAfterBuy != got {
		t.Errorf("There should be %d books left in stock after buying %d books from a stock of %d books,  got %d",
			leftInStockAfterBuy, booksBought, b.Copies, got)
	}

}
