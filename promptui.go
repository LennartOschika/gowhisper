package main

import (
	"errors"
	"github.com/chzyer/readline"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"strings"
)

type noBellStdout struct{}

func (n *noBellStdout) Write(p []byte) (int, error) {
	if len(p) == 1 && p[0] == readline.CharBell {
		return 0, nil
	}
	return readline.Stdout.Write(p)
}

func (n *noBellStdout) Close() error {
	return readline.Stdout.Close()
}

var NoBellStdout = &noBellStdout{}

func promptCommand() (int, error) {
	prompt := promptui.Select{
		Label:  "What do you want to do?",
		Items:  []string{"Transcribe video", "Change output folder", "Change API key"},
		Stdout: NoBellStdout,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return 0, err
	}
	return i, nil
}

func getFilenames() []os.DirEntry {
	directory, _ := os.Getwd()
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func promptFilename() (string, error) {
	filenames := getFilenames()

	searcher := func(input string, index int) bool {
		file := filenames[index]
		name := strings.Replace(strings.ToLower(file.Name()), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:             "Filename",
		Items:             filenames,
		Size:              4,
		Searcher:          searcher,
		StartInSearchMode: true,
		Stdout:            NoBellStdout,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return filenames[i].Name(), nil
}

func checkFile(filename string) error {
	log.Println("Filename provided: " + filename)
	if !fileExists(filename) {
		log.Println(os.Getwd())
		return errors.New("File could not be opened")
	}
	return nil
}

func promptAPIKey() error {
	prompt := promptui.Prompt{
		Label:  "API Key",
		Mask:   '*',
		Stdout: NoBellStdout,
	}

	input, err := prompt.Run()
	if err != nil {
		return err
	}
	err = SetAPIKey(input)

	return err
}

func setupPromptUI() {
	promptAPIKey()
	promptOutputfolder()
}

func promptOutputfolder() error {
	validator := promptui.ValidateFunc(func(input string) error {
		input = strings.TrimSpace(input)
		input = strings.TrimRight(input, "\\")
		validDirectory := directoryExists(input)
		if validDirectory {
			return nil
		}

		return errors.New("Not a valid directory")
	})

	prompt := promptui.Prompt{
		Label:    "Output folder",
		Stdout:   NoBellStdout,
		Validate: validator,
	}

	directory, err := prompt.Run()
	if err != nil {
		return err
	}

	err = SetOutputDirectory(directory)
	if err != nil {
		return err
	}
	log.Printf("Set output directory to %s\n", directory)
	return nil
}
