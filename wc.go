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
var fileName string

func handleError(message string, err error) {
	if err != nil {
		fmt.Println(message, ": ", err)
		os.Exit(1)
	}

}

func LineCounter(file io.Reader) int {
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++

	}
	if err := scanner.Err(); err != nil {
		if err == bufio.ErrTooLong {
			count = 0
		} else {
			fmt.Println("There was a problem counting lines:", err)
			os.Exit(1)
		}
	}

	return count
}

func ByteCounter(file io.Reader) int {
	count := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("There was a problem scanning the input: ", err)
	}
	return count
}

func init() {
	flag.BoolVar(&countBytes, "c", false, "Count bytes")
	flag.BoolVar(&countLines, "l", false, "Count Lines")
}
func main() {

	var (
		file      io.Reader
		buf       *bytes.Buffer
		lineCount int
		byteCount int
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
		file = bytes.NewReader(buf.Bytes())
		byteCount = ByteCounter(file)
		fmt.Print(byteCount, " ")
	}
	if countLines {
		file = bytes.NewReader(buf.Bytes())
		lineCount = LineCounter(file)
		fmt.Print(lineCount, " ")
	}
	fmt.Print(fileName, " ")

}
