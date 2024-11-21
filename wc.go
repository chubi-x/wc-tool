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
	"flag"
	"fmt"
	"io"
	"os"
)

var countBytes bool
var file_name string
var file io.Reader

func CountBytes(file io.Reader, writer io.Writer) {
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

	flag.Parse()
	if file_arg := flag.Arg(0); file_arg == "" {
		file = os.Stdin
	} else {
		file_name = file_arg
		open_file, err := os.Open(file_name)
		defer open_file.Close()
		file = open_file
		if err != nil {
			fmt.Println("Unable to open file: ", file_name)
			os.Exit(1)
		}
	}
	if countBytes {
		ByteCounter(file, os.Stdout)
	}
	fmt.Print(" ", file_name, " ")

}
