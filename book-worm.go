package main

import (
	"encoding/json"
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
