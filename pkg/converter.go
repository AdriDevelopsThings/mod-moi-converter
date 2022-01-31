package pkg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getModFilename(moiFilename string) string {
	return moiFilename[:len(moiFilename)-3] + "MOD"
}

func convertFFMPEG(sourceFilename string, destinationFilename string) error {
	cmd := exec.Command("ffmpeg", "-i", sourceFilename, "-c:v", "libx264", "-c:a", "aac", "-y", destinationFilename)
	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error while running command: %v: stdout: %q\n", err, stderr.String())
	}
	return nil
}

func convertModMoi(moiFilename string, destination string) error {
	moi, err := ReadMoiFile(moiFilename)
	if err != nil {
		return err
	}
	builder := new(strings.Builder)
	fmt.Fprintf(builder, filepath.FromSlash(destination+"/%04d/%02d"), moi.Year, moi.Month)
	os.MkdirAll(builder.String(), os.ModePerm)
	builder = new(strings.Builder)
	fmt.Fprintf(builder, filepath.FromSlash(destination+"/%04d/%02d/%02d_%02d_%02d_video.mp4"), moi.Year, moi.Month, moi.Day, moi.Hour, moi.Minutes)
	filename := builder.String()
	modFilename := getModFilename(moiFilename)
	return convertFFMPEG(modFilename, filename)
}

func ConvertModMoi(source_directory string, destination_directory string) error {
	moiFiles, err := FindFiles(source_directory, ".moi")
	if err != nil {
		return err
	}
	for _, f := range moiFiles {
		fmt.Printf("Convert %s\n", f)
		err := convertModMoi(f, destination_directory)
		if err != nil {
			fmt.Printf("Error while converting %s: %v\n", f, err)
		}
	}
	return nil
}
