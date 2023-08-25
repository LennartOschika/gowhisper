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

func directoryExists(dirname string) bool {
	_, err := os.Stat(dirname)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func createEnvFile() error {
	_, err := os.Create(getEnvPath())
	if err != nil {
		return err
	}
	return nil
}

func checkEnvPath() string {
	envFilePath := getEnvPath()
	if fileExists(envFilePath) {
		return envFilePath
	}
	return ""
}

func getEnvPath() string {
	exePath, _ := os.Executable()

	exeDirectory := filepath.Dir(exePath)

	envFilePath := filepath.Join(exeDirectory, ".whisperenv")
	return envFilePath

}

func LoadEnvironmentVariables() error {
	envFilePath := checkEnvPath()
	if envFilePath == "" {
		return nil
	}
	err := godotenv.Load(envFilePath)
	if err != nil {
		return err
	}
	return nil
}

func SetOutputDirectory(path string) error {
	if path == "" {
		return nil
	}
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

	envFilePath := checkEnvPath()

	currentEnv, err := godotenv.Read(envFilePath)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range newEnvVar {
		currentEnv[key] = value
	}
	err = godotenv.Write(currentEnv, envFilePath)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
