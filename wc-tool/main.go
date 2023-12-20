package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {
	countWords := flag.Bool("w", false, "bool flag to count words")
	countLines := flag.Bool("l", false, "bool flag to count lines")
	countBytes := flag.Bool("c", false, "bool flag to count bytes")
	countChars := flag.Bool("m", false, "bool flag to count chars")

	flag.Parse()

	fmt.Println("countWords:", *countWords)
	fmt.Println("countLines:", *countLines)
	fmt.Println("countBytes:", *countBytes)
	fmt.Println("countChars:", *countChars)

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
	if *countBytes {
		output += strconv.Itoa(CountBytes(content)) + "\t"
	}

	if *countLines {
		output += strconv.Itoa(CountLines(content)) + "\t"
	}

	if *countWords {
		output += strconv.Itoa(CountWords(content)) + "\t"
	}

	output += fileName
	fmt.Println(output)

}
