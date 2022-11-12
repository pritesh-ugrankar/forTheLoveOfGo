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

	wantLeftInStock := 98
	buyTheBook := bookstore.Buy(b)
	copiesBought := buyTheBook.Copies
	remainingAfterBuy := b.Copies - copiesBought

	if wantLeftInStock != remainingAfterBuy {
		t.Errorf("Want %d left in stock after buying 2 book(s) from a stock of %d books, but got %d",
			wantLeftInStock, b.Copies, copiesBought)
	}
}
