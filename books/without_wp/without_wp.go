package withoutwp

import (
	"fmt"
	"time"
)

// Represents a simple book
type Book struct {
	title string
}

func MoveBooks() {
	nOfBooks := 50

	booksInOldStore := []*Book{}
	booksInNewStore := []*Book{}

	for bookNumb := 0; bookNumb <= nOfBooks; bookNumb++ {
		booksInOldStore = append(booksInOldStore, &Book{title: fmt.Sprintf("Story %d", bookNumb)})
	}

	for _, bookInOldStore := range booksInOldStore {
		fmt.Printf("Moving book %s to new store\n", bookInOldStore.title)

		time.Sleep(time.Second * 1)
		booksInNewStore = append(booksInNewStore, bookInOldStore)

		fmt.Printf("Moved book %s to new store with success!\n", bookInOldStore.title)
	}
}
