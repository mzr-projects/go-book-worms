package main

import (
	"log"
	"sort"
)

func main() {

	log.Println("######## Load Book Worms ########")
	worms, err := loadBookWorms("testdata/books.json")
	log.Println(worms)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("######## Book Count ########")
	printBookWormsMap(bookCount(worms))

	log.Println("######## Common Books ########")
	findCommonBooks(worms)

	log.Println("######## Sorted Books By Title ########")
	var sortedBooks = getBooks(worms)
	sort.Slice(sortedBooks, func(i, j int) bool {
		if sortedBooks[i].Title < sortedBooks[j].Title {
			return true
		}
		return sortedBooks[i].Title == sortedBooks[j].Title
	})
	printBooksSlice(sortedBooks)

	log.Println("######## Sorted Books By Author ########")
	var sortedByAuthor = sortBooksByAuthor(getBooks(worms))
	printBooksSlice(sortedByAuthor)

}
