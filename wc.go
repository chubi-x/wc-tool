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
var file_name string
var file io.Reader
var fileName string

func ByteCounter(file io.Reader, writer io.Writer) {
	count := 0
	byte_slice := make([]byte, 1024)
	if file == nil {
		fmt.Println("The provided file does not exist!")
		os.Exit(1)
	}
	for {
		n, err := file.Read(byte_slice)
		count += n
		if err == io.EOF {
			break
		}
	}
	fmt.Fprint(writer, count)
}

func init() {
	flag.BoolVar(&countBytes, "c", false, "Count bytes")
}
func main() {

	var file io.Reader
	var buf *[]byte
	flag.Parse()
	if file_arg := flag.Arg(0); file_arg == "" {
		file = os.Stdin
	} else {
		fileName = file_arg
		open_file, err := os.Open(fileName)

		byteSlice, copyErr := io.ReadAll(open_file)
		buf = &byteSlice

		if err != nil || copyErr != nil {
			fmt.Println("Unable to open file: ", fileName)
			os.Exit(1)
		}
		defer open_file.Close()
	}
	if countBytes {
		file = bytes.NewReader(*buf)
		ByteCounter(file, os.Stdout)
	}
	fmt.Print(" ", file_name, " ")

}
