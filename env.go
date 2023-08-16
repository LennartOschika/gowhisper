package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnviornmentVariables() error {
	env, err := godotenv.Unmarshal("KEY2=cock")
	if err != nil {
		fmt.Println(err)
	}

	err = godotenv.Write(env, ".env")

	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func SetOutputDirectory(path string) error {
	updateString := "OUTPUTPATH=" + path
	err := updateEnv(updateString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SetAPIKey(apikey string) error {
	updateString := "APIKEY=" + apikey
	err := updateEnv(updateString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func updateEnv(variableString string) error {
	newEnvVar, err := godotenv.Unmarshal(variableString)
	if err != nil {
		fmt.Println(err)
	}
	currentEnv, err := godotenv.Read(".env")

	for key, value := range newEnvVar {
		currentEnv[key] = value
	}

	err = godotenv.Write(currentEnv, ".env")

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
