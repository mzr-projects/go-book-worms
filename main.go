package main

import "log"

func main() {
	worms, err := loadBookWorms("testdata/books.json")
	log.Println(worms)
	if err != nil {
		log.Println(err)
		return
	}
}
