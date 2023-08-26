package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func transcribe(audioFile string) error {
	if audioFile == "" {
		return errors.New("No audio file provided")
	}
	outputDir, ok := os.LookupEnv("OUTPUTDIR")
	if !ok {
		fmt.Println("Output directory is not set/found. Saving file in current directory.\n Use \"gowhisper sp -dir <directory>\" to set one.")
		outputDir, _ = os.Getwd()
	}

	c := openai.NewClient(os.Getenv("APIKEY"))

	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: audioFile,
		Language: "en",
		Format:   openai.AudioResponseFormatVTT,
	}

	fmt.Println("Transcribing video...")
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		return err
	}

	fullFilePath := filepath.Join(outputDir, audioFile)
	extension := path.Ext(fullFilePath)
	fullFilePath = strings.ReplaceAll(fullFilePath, extension, ".vtt")

	fmt.Printf("Creating file %s\n", fullFilePath)
	f, err := os.Create(fullFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(resp.Text)
	if err != nil {
		return err
	}
	fmt.Println("Successfully wrote subtitles to file.")
	return nil
}
