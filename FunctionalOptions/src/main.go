package main

import (
	"Go-Patterns/FunctionalOptions/src/file"
)

func main() {
	_, err := file.New("./empty.txt")
	if err != nil {
		panic(err)
	}
	_, err = file.New("./file.txt", file.UID(1000), file.Contents("Functional Options"))
	if err != nil {
		panic(err)
	}
}
