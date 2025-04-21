package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// read the current directory
	files, err := os.ReadDir(".")

	// if there is any error, print
	// it and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// for each file in the current
	// directory, print the file name
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
