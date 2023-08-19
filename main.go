package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func setupFirst() bool {
	if checkEnvPath() == "" {
		return true
	}
	return false
}

func setupVariables() {
	setupApp := &cli.App{
		Action: func(c *cli.Context) error {
			fmt.Print(".whisperenv file not found. \nPlease setup the following values before you can use the program.\n\n")
			reader := bufio.NewReader(os.Stdin)

			apiKey := askAPIKey(reader)

			outputPath := loopOutputPath(reader)

			fmt.Printf("Got api KEY: %s\nGot output path: %s\n", apiKey, outputPath)
			err := createEnvFile()
			if err != nil {
				return err
			}
			err = SetAPIKey(apiKey)
			if err != nil {
				return err
			}
			err = SetOutputDirectory(outputPath)
			if err != nil {
				return err
			}
			fmt.Print("Successfully saved your parameters. Run gowhisper again to see the available commands and how to use them.")
			return nil
		},
	}

	err := setupApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	err := LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
	}

	if setupFirst() {
		setupVariables()
		return
	}

	app := &cli.App{
		//Name: "gowhisper",
		Usage:     "Easily transcribe audio of video files.\n",
		UsageText: "gowhisper.exe transcribe some_file.mp3",

		Commands: cli.Commands{
			{
				Name:      "setkey",
				Aliases:   []string{"sk"},
				Usage:     "Save a new openai api key",
				UsageText: "gowhisper sekey",
				Action: func(c *cli.Context) error {
					reader := bufio.NewReader(os.Stdin)
					apiKey := askAPIKey(reader)
					err := SetAPIKey(apiKey)
					return err
				},
			},
			{
				Name:      "setPath",
				Aliases:   []string{"sp"},
				Usage:     "Save output files to a different directory",
				UsageText: "gowhisper setPath <absolute path>",
				Action: func(c *cli.Context) error {
					reader := bufio.NewReader(os.Stdin)
					outputPath := loopOutputPath(reader)
					err := SetOutputDirectory(outputPath)
					return err
				},
			},
			{
				Name:      "transcribe",
				Aliases:   []string{"t"},
				Usage:     "Transcribe audio of file file",
				UsageText: "gowhisper transcribe -f your_file.mp3",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f", "a"},
						Usage:   "Name of the file",
					},
				},
				Action: func(c *cli.Context) error {
					filename := c.String("file")
					if !fileExists(filename) {
						filename = c.Args().First()
						if !fileExists(filename) {
							return errors.New("No correct filename was provided.")
						}
					}

					fmt.Println("Filename provided: " + filename)
					err := transcribe(filename)
					return err
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
