// Package main implements the Whispercleaner application, which removes timestamps
// from Whisper transcription output.
//
// Whispercleaner reads from standard input, removes timestamp information from each line,
// and writes the cleaned text to standard output. It is designed to be used in Unix-like
// pipelines for processing Whisper transcription output.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// version is the current version of the Whispercleaner application.
const version = "1.0.0"

var (
	// timestampRegex is a regular expression used to match and remove timestamp
	// information from the beginning of each line.
	timestampRegex = regexp.MustCompile(`^\[.*?\]\s*`)

	// flags
	helpFlag    bool
	versionFlag bool
	verboseFlag bool
)

func main() {
	flag.BoolVar(&helpFlag, "help", false, "Display usage information")
	flag.BoolVar(&helpFlag, "h", false, "Display usage information (shorthand)")
	flag.BoolVar(&versionFlag, "version", false, "Display version information")
	flag.BoolVar(&versionFlag, "v", false, "Display version information (shorthand)")
	flag.BoolVar(&verboseFlag, "verbose", false, "Enable verbose logging")
	flag.Parse()

	if helpFlag {
		usage()
		return
	}

	if versionFlag {
		fmt.Printf("Whispercleaner version %s (%s/%s)\n", version, runtime.GOOS, runtime.GOARCH)
		return
	}

	setupLogging()

	if err := processInput(os.Stdin, os.Stdout); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// usage prints the usage information for the Whispercleaner application.
func usage() {
	fmt.Fprintf(os.Stderr, "Whispercleaner - Clean Whisper transcription output\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Whispercleaner removes timestamps from Whisper transcription output.\n")
	fmt.Fprintf(os.Stderr, "It reads from standard input and writes to standard output.\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExample usage:\n")
	fmt.Fprintf(os.Stderr, "  cat whisper_output.txt | %s > cleaned_output.txt\n", os.Args[0])
}

// setupLogging configures the logging based on the verbose flag.
func setupLogging() {
	if verboseFlag {
		log.SetOutput(os.Stderr)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	} else {
		log.SetOutput(io.Discard)
	}
}

// processInput reads input from the provided io.Reader, processes each line
// to remove timestamps, and writes the cleaned output to the provided io.Writer.
// It returns an error if there are any issues reading from the input or writing to the output.
func processInput(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriter(w)
	defer func() {
		if err := writer.Flush(); err != nil {
			log.Printf("Error flushing writer: %v", err)
		}
	}()

	var lineCount, cleanedCount int

	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		cleanedLine := cleanLine(line)
		if cleanedLine != "" {
			cleanedCount++
			if _, err := writer.WriteString(cleanedLine + "\n"); err != nil {
				return fmt.Errorf("error writing output on line %d: %w", lineCount, err)
			}
		}
		log.Printf("Processed line %d", lineCount)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	log.Printf("Processed %d lines, cleaned %d lines", lineCount, cleanedCount)
	return nil
}

// cleanLine removes the timestamp from the beginning of the provided line
// and trims any leading or trailing whitespace.
// It returns the cleaned line as a string.
func cleanLine(line string) string {
	return strings.TrimSpace(timestampRegex.ReplaceAllString(line, ""))
}
