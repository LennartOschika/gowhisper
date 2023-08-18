package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name: "gowhisper",
		Usage: "Easily convert audio files into audio.\n" +
			"Run \"gowhisper i\" to install the program and see to add it to your system's path variable." +
			"That way you can call this program from anywhere.\n" +
			"This will also create a .whisperenv file which stores your credentials and output directory.",
		Commands: cli.Commands{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Creates .env file and shows how to add a path variable.",
				Action: func(c *cli.Context) error {
					executablePath, err := os.Executable()
					if err != nil {
						return cli.Exit("Could not get current executable path.", 1)
					}

					programDir := filepath.Dir(executablePath)
					fmt.Printf("To add the program directory to PATH, run the following command the following to your path environment variable\n")
					fmt.Printf("%s\n", programDir)
					envFilePath := programDir + "\\.whisperenv"
					_, err = os.Stat(envFilePath)
					if err != nil {
						file, err := os.Create(programDir + "\\.whisperenv")
						if err != nil {
							return cli.Exit("Could not create .whisperenv file", 1)
						} else {
							_ = file.Close()
							return cli.Exit("Created .whisperenv file", 0)
						}
					} else {
						return cli.Exit(".whisperenv file already exists. Did not overwrite it.", 0)
					}
				},
			},
			{
				Name:    "setkey",
				Aliases: []string{"sk"},
				Usage:   "Saves the openai API Key.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "key",
						Aliases: []string{"k"},
						Usage:   "Your openai key",
					},
				},
				Action: func(c *cli.Context) error {
					apiKey := c.String("key")
					if apiKey == "" {
						return cli.Exit("Please provide an API key.", 1)
					}
					err := SetAPIKey(apiKey)
					if err != nil {
						return cli.Exit("Something failed saving API key.", 1)
					}
					fmt.Printf("Set API Key: %s\n", apiKey)
					return nil
				},
			},
			{
				Name:    "setPath",
				Aliases: []string{"sp"},
				Usage:   "Sets the full path to the output directory",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "folder",
						Aliases: []string{"p", "f", "dir"},
						Usage:   "The path to the output folder",
					},
				},
				Action: func(c *cli.Context) error {
					path := c.String("folder")
					if path == "" {
						return cli.Exit("Please provide an output directory.", 1)
					}
					path = strings.TrimRight(path, "\\")

					_, err := os.Stat(path)
					if err != nil {
						if os.IsNotExist(err) {
							return cli.Exit("Output directory does not exist.", 1)
						}
					}
					err = SetOutputDirectory(path)
					if err != nil {
						return cli.Exit("Something failed saving output path.", 1)
					}
					fmt.Printf("Set output directory: %s\n", path)
					return nil
				},
			},
			{
				Name:    "transcribe",
				Aliases: []string{"t"},
				Usage:   "Test working directory",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f", "a"},
						Usage:   "Name of the audio file",
					},
				},
				Action: func(c *cli.Context) error {
					filename := c.String("file")
					fmt.Println("Filename provided: " + filename)
					transcribe(filename)
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
