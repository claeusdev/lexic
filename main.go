package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// get started here

	// define boolean flag -l to count lines instead of words

	lines := flag.Bool("l", false, "Count lines")
	totalBytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

	// count the number of words received from stdin
	// print to console

	fmt.Println(count(os.Stdin, *lines, *totalBytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// read text from Reader
	scanner := bufio.NewScanner(r)

	// split by words insteead of by lines

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// start counting
	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
