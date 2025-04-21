package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Retrieve information about the 'appsettings.json' file in the current directory
	_, err := os.Stat("appsettings.json")

	// If there is an error in retrieving the file's info...
	if err != nil {
		// Check if it's that the file does not exist
		if os.IsNotExist(err) {
			fmt.Println("Error: The required configuration file 'appsettings.json' was not found in the directory of the executed program.")
			os.Exit(1)
		} else {
			// Else, print the error and exit with code 1
			log.Fatal(err)
		}
	}

	// The file exists!
	fmt.Println("The appsettings.json file exists")
}
