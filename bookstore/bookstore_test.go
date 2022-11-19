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
	if err != nil {
		t.Fatalf("Fatal Error!! %v", err)
	}

	copiesBought := buyTheBooks.Copies

	stockDiff := b.Copies - leftInStockAfterBuy

	if leftInStockAfterBuy != copiesBought {

		t.Errorf("There should be %d books left in stock after buying %d books from a stock of %d books,  but got %d",
			leftInStockAfterBuy, stockDiff, b.Copies, copiesBought)
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
