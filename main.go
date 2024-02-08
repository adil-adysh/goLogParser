package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var showTimestamp bool

func init() {
	// Define command-line flag to show timestamp
	flag.BoolVar(&showTimestamp, "show-timestamp", false, "Show timestamp in log messages")
}

func plainLog(logContent string) []string {
	var messages []string

	// Split the log content into lines
	lines := strings.Split(logContent, "\n")

	// Iterate through each line and extract the message
	for _, line := range lines {
		// Split the line into segments separated by "|"
		segments := strings.Split(line, "|")

		// Extract the message from the last segment
		message := strings.TrimSpace(segments[len(segments)-1])

		// Hide timestamp if showTimestamp flag is not set
		if !showTimestamp {
			// Use regex to remove timestamp if present
			re := regexp.MustCompile(`\d{4}/\d{2}/\d{2} (\d{2}:\d{2}:\d{2}) `)
			message = re.ReplaceAllString(message, "")
		} else {
			// Format timestamp to HH:MM:SS if showTimestamp flag is set
			re := regexp.MustCompile(`(\d{4}/\d{2}/\d{2} )(\d{2}:\d{2}:\d{2})`)
			message = re.ReplaceAllString(message, "$2")
		}

		// Append the message to the messages slice
		messages = append(messages, message)
	}

	return messages
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// Read input from stdin
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	logContent := string(inputBytes)

	// Extract the messages using the plainLog function
	messages := plainLog(logContent)

	// Print the extracted messages
	for _, message := range messages {
		fmt.Println(message)
	}
}
