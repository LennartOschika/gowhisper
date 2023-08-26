package main

import (
	"log"
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

	method := surveyWhatToDo()
	switch method {
	case 0:
		err = surveyTranscribe()
		if err != nil {
			log.Fatal(err)
		}
	case 1:
		err := surveyOutputfolder()
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		err = surveyAPIKey()
		if err != nil {
			log.Fatal(err)
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
