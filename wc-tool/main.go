package main

import (
	"flag"
	"fmt"
	"os"
)


func main() {
	countWords := flag.Bool("w", false, "bool flag to count words")
	countLines := flag.Bool("l", false, "bool flag to count lines")
	countBytes := flag.Bool("m", false, "bool flag to count bytes")
	countChars := flag.Bool("c", false, "bool flag to count chars")

	flag.Parse()

	fmt.Println("countWords:", *countWords)
	fmt.Println("countLines:", *countLines)
	fmt.Println("countBytes:", *countBytes)
	fmt.Println("countChars:", *countChars)

	extraArgs := flag.Args()

	var fileName string
	if len(extraArgs) == 1 {
		fileName = extraArgs[0]
		fmt.Println("fileName:", fileName)
	} else {
		panic("No file name provided :(")
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

}
