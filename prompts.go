package main

import (
	"bufio"
	"fmt"
	"strings"
)

func askAPIKey(reader *bufio.Reader) string {
	var apiKey string
	validKey := false
	for !validKey {
		fmt.Print("Enter your api key: ")
		apiKey, _ = reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		if apiKey != "" {
			validKey = true
		} else {
			fmt.Println("The api key may not be empty!")
		}
	}
	return apiKey
}

func askOutputPath(reader *bufio.Reader) string {
	fmt.Print("Enter the desired output path (optional): ")
	outputPath, _ := reader.ReadString('\n')
	outputPath = strings.TrimSpace(outputPath)

	return outputPath
}

func loopOutputPath(reader *bufio.Reader) string {
	validDir := false
	var outputPath string
	for !validDir {
		outputPath = askOutputPath(reader)
		//When not empty it has to be valid
		if outputPath != "" {
			if !directoryExists(outputPath) {
				fmt.Println("Directory does not exist")
			} else {
				fmt.Println("Directory exists")
				validDir = true
			}
		} else {
			validDir = true
		}
	}
	return outputPath
}
