package main

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"strings"
)

func surveyWhatToDo() int {
	var whatToDo int

	prompt := &survey.Select{
		Message: "What do you want to do:",
		Options: []string{"Transcribe video", "Change output folder", "Change API key"},
	}
	survey.AskOne(prompt, &whatToDo)
	return whatToDo
}

func surveyOutputfolder() error {
	path := ""
	//outputSelect := &survey.Input{
	//	Message: "Output directory:",
	//}

	outputSelect := []*survey.Question{
		{
			Prompt: &survey.Input{Message: "Output directory:"},
			Validate: func(ans interface{}) error {
				ok := directoryExists(ans.(string))
				if ok {
					return nil
				} else {
					return errors.New("Output directory does not exist.")
				}
			},
		},
	}

	survey.Ask(outputSelect, &path)
	return SetOutputDirectory(path)
}

func surveyAPIKey() error {
	apiKey := ""
	keySelect := &survey.Input{
		Message: "API key:",
	}

	survey.AskOne(keySelect, &apiKey)
	return SetAPIKey(apiKey)
}

func surveyTranscribe() error {
	prompt2 := &survey.Select{
		Message: "File to transcribe: ",
		Options: getFilenamesString(),
		Filter: func(filter string, value string, index int) bool {
			value = strings.Replace(strings.ToLower(value), " ", "", -1)
			filter = strings.Replace(strings.ToLower(filter), " ", "", -1)

			return strings.Contains(value, filter)
		},
	}

	file := ""
	survey.AskOne(prompt2, &file)
	return transcribe(file)
}

func getFilenamesString() []string {
	wd, _ := os.Getwd()
	files, _ := os.ReadDir(wd)

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}
