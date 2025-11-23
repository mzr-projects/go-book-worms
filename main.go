package main

import (
	"log"
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
}
