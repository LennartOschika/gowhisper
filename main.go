package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func setupFirst() bool {
	if checkEnvPath() == "" {
		return true
	}
	return false
}

func main() {

	var err error

	if setupFirst() {

		err = createEnvFile()
		if err != nil {
			log.Fatal(err)
		}
		setup()

	}

	err = LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
	}

	checkEnvVarsSet()

	method := surveyWhatToDo()
	for _, entry := range functions {
		if entry.Id == method {
			err = entry.Function()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func setup() {

	err := surveyOutputfolder()
	if err != nil {
		log.Fatal(err)
	}

	err = surveyAPIKey()
	if err != nil {
		log.Fatal(err)
	}
}

func checkEnvVarsSet() {
	value, exists := os.LookupEnv("APIKEY")
	if !exists || strings.TrimSpace(value) == "" {
		fmt.Println(".whisperenv file exists but no API Key was found")
		err := surveyAPIKey()
		if err != nil {
			log.Println(err)
		}
	}
	value, exists = os.LookupEnv("OUTPUTDIR")
	if !exists || strings.TrimSpace(value) == "" {
		fmt.Println("No output directory set. Using current directory instead.")
	}

	err := LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
	}

}
