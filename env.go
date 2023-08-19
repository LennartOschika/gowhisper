package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getEnvPath() string {
	exePath, _ := os.Executable()

	exeDirectory := filepath.Dir(exePath)

	log.Println(exeDirectory)

	envFilePath := filepath.Join(exeDirectory, ".whisperenv")
	if fileExists(envFilePath) {
		log.Println("Environment variable file found")
	} else {
		_, err := os.Create(envFilePath)
		if err != nil {
			return ""
		}
		log.Println("Created new environment file in executable directory")

	}
	return envFilePath
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
