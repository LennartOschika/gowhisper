package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func getEnvPath() string {
	exePath, _ := os.Executable()

	exeDirectory := filepath.Dir(exePath)

	envFilePath := ".whisperenv"
	_, err := os.Stat(envFilePath)
	if err == nil {
		return envFilePath
	}

	envFilePath = filepath.Join(exeDirectory, ".whisperenv")
	_, err = os.Stat(envFilePath)
	if err == nil {
		return envFilePath
	}

	return ""

}

func LoadEnvironmentVariables() error {

	envFilePath := getEnvPath()
	err := godotenv.Load(envFilePath)
	if err != nil {
		return err
	}
	return nil
}

func SetOutputDirectory(path string) error {
	updateString := "OUTPUTDIR=" + path
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

	envFilePath := getEnvPath()

	currentEnv, err := godotenv.Read(envFilePath)

	for key, value := range newEnvVar {
		currentEnv[key] = value
	}

	err = godotenv.Write(currentEnv, envFilePath)

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
