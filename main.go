package main

import (
	"fmt"
	"os"

	"github.com/adridevelopsthings/mod-moi-converter/pkg"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("mod-moi-converter", "Convert mod videos with moi informations to mp4 files.")
	source_directory := parser.String("s", "source-directory", &argparse.Options{Required: true, Help: "Directory with mod and moi files."})
	destination_directory := parser.String("d", "destination-directory", &argparse.Options{Required: true, Help: "Directory with new mp4 files."})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Printf("Error while parsing args %v\n", err)
		return
	}
	err = pkg.ConvertModMoi(*source_directory, *destination_directory)
	if err != nil {
		fmt.Printf("Error while converting: %v\n", err)
	} else {
		fmt.Println("Success.")
	}
}
