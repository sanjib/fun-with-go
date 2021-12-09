package main

// The IterableCollection interface defines the createIterator
// function, which returns an iterator object
type IterableCollection interface {
	createIterator() iterator
}

// The iterator interface contains the hasNext and next functions
// which allow the collection to return items as needed
type iterator interface {
	hasNext() bool
	next() *Book
}

// BookIterator is a concrete iterator for a Book collection
type BookIterator struct {
	current int
	books   []Book
}

func (b *BookIterator) hasNext() bool {
	// implement hasNext()
	if b.current < len(b.books) {
		return true
	}
	return false
}

func (b *BookIterator) next() *Book {
	// implement next()
	if b.hasNext() {
		book := b.books[b.current]
		b.current++
		return &book
	}
	return nil
}
