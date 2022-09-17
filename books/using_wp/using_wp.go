package usingwp

import (
	"fmt"
	"time"
)

// Represents a simple book
type Book struct {
	title string
}

// Friend are our worker
type Friend struct {
	name string
}

// newFriend must create our worker
func newFriend(name string) *Friend {
	return &Friend{
		name: name,
	}
}

func MoveBooks() {
	nOfBooks := 50

	booksInOldStore := make(chan *Book, nOfBooks)
	booksInNewStore := make(chan *Book)

	friends := createFriends()

	// there will be 5 workers in pool
	for _, friend := range friends {
		go friend.moveBook(booksInOldStore, booksInNewStore)
	}

	for bookNumb := 0; bookNumb <= nOfBooks; bookNumb++ {
		booksInOldStore <- &Book{title: fmt.Sprintf("Story %d", bookNumb)}
	}
	close(booksInOldStore)

	for i := 0; i <= nOfBooks; i++ {
		<-booksInNewStore
	}
}

// moveBook is our work function
func (f *Friend) moveBook(booksInOldStore <-chan *Book, booksInNewStore chan<- *Book) {
	for bookInOldStore := range booksInOldStore {
		fmt.Printf("%s are moving book %s to new store\n", f.name, bookInOldStore.title)

		time.Sleep(time.Second * 1)

		fmt.Printf("%s moved book %s to new store with success!\n", f.name, bookInOldStore.title)

		booksInNewStore <- bookInOldStore
	}
}

// createFriends will return a slice of Friend (the pool of workers)
func createFriends() []*Friend {
	var friends []*Friend

	fred := newFriend("Fred")
	daphne := newFriend("Daphne")
	saggy := newFriend("Saggy")
	velma := newFriend("Velma")
	scooby := newFriend("Scooby")

	return append(friends, fred, daphne, saggy, velma, scooby)
}
