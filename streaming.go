package main

import (
	"os/exec"
	"log"
)

func processVideoForFastStart(filePath string) (string, error) {
	outputFilepath := filePath + ".processing"

	cmd := exec.Command("ffmpeg",
		"-i", filePath, "-c",
		"copy", "-movflags",
		"faststart", "-f", "mp4",
		outputFilepath,
	)

    if err := cmd.Run(); err != nil {
        log.Print("Couldn't run cmd (processVideo): ", err)
        return "", err
    }

	return outputFilepath, nil
}
