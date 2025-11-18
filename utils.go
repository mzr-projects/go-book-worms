package main

import "fmt"

func PrintLog(message string) {
	fmt.Println(message)
}

func PrintError(err error) {
	fmt.Println(err.Error())
}
