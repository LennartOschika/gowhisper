package main

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"log"
	"os"
	"strings"
)

var functions = []struct {
	Id        int
	Name      string
	Function  func() error
	EnvVar    string
	Obfuscate bool
}{
	{Id: 0, Name: "Transcribe video", Function: surveyTranscribe},
	{Id: 1, Name: "Change output folder", Function: surveyOutputfolder, EnvVar: "OUTPUTDIR"},
	{Id: 2, Name: "Change API key", Function: surveyAPIKey, EnvVar: "APIKEY", Obfuscate: true},
}

func getCurrentValue(variable string, obfuscate bool) string {
	if variable == "" {
		return ""
	}
	value := os.Getenv(variable)
	if obfuscate {
		value = "sk-****" + value[len(value)-5:]
	}
	return "Current: " + value
}

func getOptions() []string {
	functionNames := make([]string, len(functions))
	for index, fn := range functions {
		functionNames[index] = fn.Name
	}
	return functionNames
}

func surveyWhatToDo() int {
	var whatToDo int
	prompt := &survey.Select{
		Message: "What do you want to do:",
		Options: getOptions(),
		Description: func(value string, index int) string {
			for _, entry := range functions {
				if entry.Id == index {
					return getCurrentValue(entry.EnvVar, entry.Obfuscate)
				}
			}
			return ""
		},
	}
	err := survey.AskOne(prompt, &whatToDo)
	log.Println(err)
	return whatToDo
}

func surveyOutputfolder() error {
	path := ""

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
	prompt := &survey.Select{
		Message: "File to transcribe: ",
		Options: getFilenamesString(),
		Filter: func(filter string, value string, index int) bool {
			value = strings.Replace(strings.ToLower(value), " ", "", -1)
			filter = strings.Replace(strings.ToLower(filter), " ", "", -1)

			return strings.Contains(value, filter)
		},
	}

	file := ""
	survey.AskOne(prompt, &file)
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
