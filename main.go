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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
