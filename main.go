package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	lines := flag.Bool("l", false, "Count lines")
	totalBytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *totalBytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
