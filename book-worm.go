package main

import "os"

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
		return nil, err
	}

	/*
		The defer statement defers the execution of a function until the surrounding function returns.
		This defer will execute after the function returns and closes the file.
	*/
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	return nil, nil
}
