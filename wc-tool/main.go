package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func CountBytes(input []byte) int {
	return len(input)
}

func CountLines(input []byte) int {
	count := 0
	for _, char := range input {
		if char == '\n' {
			count++
		}
	}
	return count
}

func CountWords(input []byte) int {
	split := strings.Fields(string(input))
	return len(split)
}

func CountChars(input []byte) int {
	nonSpecialCharCount := utf8.RuneCountInString(string(input))
	return nonSpecialCharCount
}

func main() {
	countWords := flag.Bool("w", false, "bool flag to count words")
	countLines := flag.Bool("l", false, "bool flag to count lines")
	countBytes := flag.Bool("c", false, "bool flag to count bytes")
	countChars := flag.Bool("m", false, "bool flag to count chars")

	flag.Parse()

	extraArgs := flag.Args()

	var fileName string
	if len(extraArgs) == 1 {
		fileName = extraArgs[0]
	} else {
		panic("No file name provided :(")
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var output string = "\t"

	if !*countWords && !*countLines && !*countBytes && !*countChars {
		*countBytes = true
		*countLines = true
		*countWords = true
	}

	if *countLines {
		output += strconv.Itoa(CountLines(content)) + "\t"
	}

	if *countWords {
		output += strconv.Itoa(CountWords(content)) + "\t"
	}

	if *countBytes {
		output += strconv.Itoa(CountBytes(content)) + "\t"
	}

	if *countChars {
		output += strconv.Itoa(CountChars(content)) + "\t"
	}

	output += fileName
	fmt.Println(output)

}
