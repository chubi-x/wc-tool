package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	countBytes bool
	countLines bool
	countWords bool
	countChars bool
	fileName   string
)

func handleError(counterType string, err error) {
	if err != nil {
		fmt.Sprintf("There was a problem counting %s: %s", counterType, err)
		os.Exit(1)
	}

}
func counter(buf *bytes.Buffer, count_object string) int {
	count := 0
	scanner := bufio.NewScanner(bytes.NewReader(buf.Bytes()))
	switch count_object {
	case "words":
		scanner.Split(bufio.ScanWords)
	case "lines":
		scanner.Split(bufio.ScanLines)
	case "bytes":
		scanner.Split(bufio.ScanBytes)
	case "chars":
		scanner.Split(bufio.ScanRunes)
	}
	for scanner.Scan() {
		count++
	}
	handleError(count_object, scanner.Err())
	return count
}
func LineCounter(buf *bytes.Buffer) int {
	return counter(buf, "lines")
}
func ByteCounter(buf *bytes.Buffer) int {
	return counter(buf, "bytes")
}
func WordCounter(buf *bytes.Buffer) int {
	return counter(buf, "words")
}
func CharacterCounter(buf *bytes.Buffer) int {
	return counter(buf, "chars")
}
func init() {
	flag.BoolVar(&countBytes, "c", false, "Count bytes")
	flag.BoolVar(&countLines, "l", false, "Count Lines")
	flag.BoolVar(&countWords, "w", false, "Count Words")
	flag.BoolVar(&countChars, "m", false, "Count Characters")
}
func main() {

	var (
		buf *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	)

	flag.Parse()

	if file_arg := flag.Arg(0); file_arg == "" {
		_, err := io.Copy(buf, os.Stdin)
		handleError("Unable to read from Stdin", err)
	} else {
		fileName = file_arg
		open_file, err := os.Open(fileName)

		_, copyErr := io.Copy(buf, open_file)
		handleError("Unable to read file: ", err)
		handleError("Error opening file: "+fileName, copyErr)
		defer open_file.Close()
	}
	if !countBytes && !countLines && !countWords && !countChars {
		countBytes, countWords, countLines = true, true, true
	}
	if countBytes {
		fmt.Print(ByteCounter(buf), " ")
	}
	if countLines {
		fmt.Print(LineCounter(buf), " ")
	}
	if countWords {
		fmt.Print(WordCounter(buf), " ")
	}
	if countChars {
		fmt.Print(CharacterCounter(buf), " ")
	}
	fmt.Print(fileName, " ")

}
