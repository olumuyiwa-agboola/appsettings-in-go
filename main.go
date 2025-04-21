package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Read the 'appsettings.json' file in the current directory
	file, err := os.ReadFile("appsettings.json")

	// If there is an error in retrieving the file's info...
	if err != nil {
		// Check if it's that the file does not exist
		if os.IsNotExist(err) {
			fmt.Println("Error: The required configuration file 'appsettings.json' was not found in the directory of the executed program.")
			os.Exit(1)
		} else {
			// Else, something else happend: print it and exit with code 1
			log.Fatal(err)
		}
	}

	var config AppSettings

	// Parse the JSON-encoded data in the file and store it in
	// the config variable, which represents the appsettings.json file
	err = json.Unmarshal(file, &config)

	// If there's any error, print it and exit with code 1
	if err != nil {
		log.Fatal(err)
	}

	// Else, print the values in the appsettings.json file
	fmt.Println(config.ConnectionStrings)
	fmt.Println(config.ConnectionStrings.SampleDb1)
	fmt.Println(config.ConnectionStrings.SampleDb2)

	fmt.Println(config.SomeOtherSettings)
	fmt.Println(config.SomeOtherSettings.Setting1)
	fmt.Println(config.SomeOtherSettings.Setting2)
}

type AppSettings struct {
	ConnectionStrings ConnectionStrings `json:"ConnectionStrings"`
	SomeOtherSettings SomeOtherSettings `json:"SomeOtherSettings"`
}

type ConnectionStrings struct {
	SampleDb1 string `json:"SampleDb1"`
	SampleDb2 string `json:"SampleDb2"`
}

type SomeOtherSettings struct {
	Setting1 string `json:"Setting1"`
	Setting2 string `json:"Setting2"`
}
