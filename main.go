package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	//err := LoadEnviornmentVariables()
	//if err != nil {
	//	log.Fatal(err)
	//}

	app := &cli.App{
		Name:  "Test name",
		Usage: "Test usage",
		Commands: cli.Commands{
			{
				Name:    "setkey",
				Aliases: []string{"sk"},
				Usage:   "Sets the Whisper API Key.",
				Action: func(c *cli.Context) error {
					apiKey := c.Args().First()
					if apiKey == "" {
						return cli.Exit("Please provide an API key.", 1)
					}
					err := SetAPIKey(apiKey)
					if err != nil {
						return cli.Exit("Something failed saving API key.", 1)
					}
					fmt.Printf("Set PAI Key: %s\n", apiKey)
					return nil
				},
			},
			{
				Name:    "setPath",
				Aliases: []string{"sp"},
				Usage:   "Sets the full path to the output directory",
				Action: func(c *cli.Context) error {
					path := c.Args().First()
					if path == "" {
						return cli.Exit("Please provide an output directory.", 1)
					}

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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
