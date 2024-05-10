package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	prev "github.com/inventor500/regex-preview"
)

func main() {
	os.Exit(mainFunc())
}

func parseArgs() (prev.OutputSettings, string) {
	fore := flag.Int("foreground", 255, "The foreground color used in matching strings")
	back := flag.Int("background", 196, "The background color used in matching strings")
	file := flag.String("f", "", "File to read sample text from")
	flag.Parse()
	sample := ""
	if *file == "" && len(flag.Args()) < 1 {
		log.Fatal("No sample data provided!")
	} else if *file == "" {
		sample = strings.Join(flag.Args(), " ")
	} else {
		opened, err := os.Open(*file)
		if err != nil {
			log.Fatalf("Unable to open sample file: %s", err)
		}
		defer opened.Close()
		_s, err := io.ReadAll(opened)
		if err != nil {
			log.Fatalf("Unable to read sample file: %s", err)
		}
		sample = string(_s)
	}
	return prev.OutputSettings{
		FColor: *fore,
		BColor: *back,
	}, sample
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// Remove the line ending
	return strings.Split(line, "\n")[0], nil
}

func mainFunc() int {
	log.SetFlags(0) // Disable logging of time
	settings, sample := parseArgs()
	for {
		input, err := readInput()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// User pressed Ctrl+d
				return 0
			} else {
				// Something else happened
				fmt.Println(err)
				return 1
			}
		}
		if input == "" {
			continue
		}
		regex, err := regexp.Compile(input)
		if err != nil {
			continue
		} else {
			fmt.Println(prev.RenderOutput(sample, regex, settings))
		}
	}
}
