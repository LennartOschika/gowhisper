package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func transcribe(audioFile string) error {
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
		fmt.Println(err)
		return cli.Exit("Could not create Transcription. See error above", 1)
	}

	fullFilePath := filepath.Join(outputDir, audioFile)
	extension := path.Ext(fullFilePath)
	fullFilePath = strings.ReplaceAll(fullFilePath, extension, ".vtt")

	fmt.Printf("Creating file %s\n", fullFilePath)
	f, err := os.Create(fullFilePath)
	if err != nil {
		return cli.Exit("Could not create outputfile "+fullFilePath, 1)
	}
	defer f.Close()
	_, err = f.WriteString(resp.Text)
	if err != nil {
		return err
	}
	fmt.Println("Successfully wrote subtitles to file.")
	return nil
}
