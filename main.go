package main

import (
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

func newInput() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			option, err := promptCommand()
			if err != nil {
				return err
			}

			switch option {
			case 0:
				filename, err := promptFilename()
				if err != nil {
					return err
				}
				err = checkFile(filename)
				if err != nil {
					return err
				}
				return transcribe(filename)
			case 1:
				err := promptOutputfolder()
				if err != nil {
					return err
				}

			case 2:
				err := promptAPIKey()
				if err != nil {
					return err
				}
				log.Println("Successfully set API key to new key.")
			}

			return nil
		},
	}
	err := app.Run(os.Args)
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
		//setupVariables()

		err := createEnvFile()
		if err != nil {
			log.Fatal(err)
		}
		setupPromptUI()
	}

	newInput()
}
