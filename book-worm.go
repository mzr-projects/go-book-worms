package main

import (
	"encoding/json"
	"log"
	"os"
)

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type BookWorm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func loadBookWorms(filePath string) ([]BookWorm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		PrintError(err)
		return nil, err
	}

	/*
		The defer statement defers the execution of a function until the surrounding function returns.
		This defer will execute after the function returns and closes the file.
	*/
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			PrintError(err)
		}
	}(f)

	/*
		Here we define a slice of BookWorm structs.
	*/
	var bookWorms []BookWorm
	/*
		Here we create a decoder and decode the JSON data into the slice of BookWorm structs.
	*/
	err = json.NewDecoder(f).Decode(&bookWorms)
	if err != nil {
		PrintError(err)
		return nil, err
	}

	return bookWorms, nil
}

func findBookWormByName(bookWorms []BookWorm, name string) *BookWorm {
	for _, bookWorm := range bookWorms {
		if bookWorm.Name == name {
			return &bookWorm
		}
	}
	return nil
}

/*
This function returns a map of books and their counts.
*/
func bookCount(bookWorms []BookWorm) map[Book]uint {
	/*
		The key of the map must be hashable, so we use Book as the key type.
		So e.g slices are not allowed as keys because they are not hashable.
	*/
	count := make(map[Book]uint)

	for _, bookWorm := range bookWorms {
		for _, book := range bookWorm.Books {
			/*
				This will increase the count of the book (related to the book key) in the map.
			*/
			count[book]++
		}
	}

	return count
}

func printBookWormsMap(bookWorms map[Book]uint) {
	for book, count := range bookWorms {
		log.Printf("%s: %d\n", book.Title, count)
	}
}

func printBooksSlice(books []Book) {
	for _, book := range books {
		log.Printf("%s: %v\n", book.Title, book.Author)
	}
}

func findCommonBooks(bookWorms []BookWorm) []Book {
	books := bookCount(bookWorms)
	var commonBooks []Book

	for book, count := range books {
		if count > 1 {
			/*
				the append method always returns a new slice, so we don't need to worry about mutating the original slice.
			*/
			commonBooks = append(commonBooks, book)
		}
	}

	printBooksSlice(commonBooks)

	return commonBooks
}
