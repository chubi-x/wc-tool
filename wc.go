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

func LineCounter(file io.Reader, writer io.Writer) {
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("There was a problem counting lines:", err)
		os.Exit(1)
	}

	fmt.Fprint(writer, " ", count)
}

func ByteCounter(file io.Reader, writer io.Writer) {
	count := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	if file == nil {
		fmt.Println("The provided file does not exist!")
		os.Exit(1)
	}
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("There was a problem scanning the input: ", err)
	}

	fmt.Fprint(writer, count)
}

func init() {
	flag.BoolVar(&countBytes, "c", false, "Count bytes")
	flag.BoolVar(&countLines, "l", false, "Count Lines")
}
func main() {

	var file io.Reader
	var buf *bytes.Buffer
	flag.Parse()

	if file_arg := flag.Arg(0); file_arg == "" {
		buf = bytes.NewBuffer(make([]byte, 0))
		_, err := io.Copy(buf, os.Stdin)
	} else {
		fileName = file_arg
		open_file, err := os.Open(fileName)

		if err != nil || copyErr != nil {
			fmt.Println("Unable to open file: ", fileName)
			os.Exit(1)
		}
		buf = bytes.NewBuffer(make([]byte, 0))
		_, copyErr := io.Copy(buf, open_file)
		defer open_file.Close()
	}
	if countBytes {
		file = bytes.NewReader(buf.Bytes())
		ByteCounter(file, os.Stdout)
	}
	if countLines {
		file = bytes.NewReader(buf.Bytes())
		LineCounter(file, os.Stdout)
	}
	fmt.Print(" ", fileName, " ")

}
