// accept command line args
// positional argument should be the last
// when its absent read from stdin
//  -c for number of bytes
//  -l for number of lines
//  -m for nhumber of characters
//  -w for number of words
//  everything is disabled by default. only enable everything when none of them is enabled.
//  zero args should call all 4 together
//  print a table showing each count

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

var countBytes bool
var countLines bool
var countWords bool
var fileName string

func handleError(counterType string, err error) {
	if err != nil {
		fmt.Sprintf("There was a problem counting %s: %s", counterType, err)
		os.Exit(1)
	}

}
func Counter(buf *bytes.Buffer, count_object string) int {
	count := 0
	scanner := bufio.NewScanner(bytes.NewReader(buf.Bytes()))
	switch count_object {
	case "words":
		scanner.Split(bufio.ScanWords)

	case "lines":
		scanner.Split(bufio.ScanLines)
	case "bytes":
		scanner.Split(bufio.ScanBytes)
	}
	for scanner.Scan() {
		count++
	}
	handleError(count_object, scanner.Err())
	return count
}

func init() {
	flag.BoolVar(&countBytes, "c", false, "Count bytes")
	flag.BoolVar(&countLines, "l", false, "Count Lines")
	flag.BoolVar(&countWords, "w", false, "Count Words")
}
func main() {

	var (
		buf *bytes.Buffer
	)

	flag.Parse()

	if file_arg := flag.Arg(0); file_arg == "" {
		buf = bytes.NewBuffer(make([]byte, 0))
		_, err := io.Copy(buf, os.Stdin)
		handleError("Unable to read from Stdin", err)
	} else {
		fileName = file_arg
		open_file, err := os.Open(fileName)

		buf = bytes.NewBuffer(make([]byte, 0))
		_, copyErr := io.Copy(buf, open_file)
		handleError("Unable to read from stdin: ", err)
		handleError("Error opening file: "+fileName, copyErr)
		defer open_file.Close()
	}
	if countBytes {
		fmt.Print(Counter(buf, "bytes"), " ")
	}
	if countLines {
		fmt.Print(Counter(buf, "lines"), " ")
	}
	if countWords {
		fmt.Print(Counter(buf, "words"), " ")
	}
	fmt.Print(fileName, " ")

}
