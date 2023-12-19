package main

import "fmt"
import (
	"flag"
	"fmt"
)


func main() {
	fmt.Println("Hello, World!")
	countWords := flag.Bool("w", false, "bool flag to count words")
	countLines := flag.Bool("l", false, "bool flag to count lines")
	countBytes := flag.Bool("m", false, "bool flag to count bytes")
	countChars := flag.Bool("c", false, "bool flag to count chars")

	flag.Parse()

	fmt.Println("countWords:", *countWords)
	fmt.Println("countLines:", *countLines)
	fmt.Println("countBytes:", *countBytes)
	fmt.Println("countChars:", *countChars)

}
