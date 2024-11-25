//what counts as whitespace characterr?
// test with spaces and tabs

package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCountBytes(t *testing.T) {
	reader := bytes.NewReader(make([]byte, 10))
	readerSize := reader.Size()
	count := ByteCounter(reader)

	if count != int(readerSize) {
		t.Errorf("Expected byte count %d, got count %d", readerSize, count)
	}
}
func TestCountLines(t *testing.T) {
	reader := strings.NewReader("This is one line \n this is another line \n this is a third line")
	lineCount := 3
	count := LineCounter(reader)

	if count != int(lineCount) {
		t.Errorf("Expected line count %d, got count %d", lineCount, count)
	}
}
func TestCountWords(t *testing.T) {
	reader := strings.NewReader("There are four words")
	wordCount := 4
	count := WordCounter(reader)

	if count != int(wordCount) {
		t.Errorf("Expected word count %d, got count %d", wordCount, count)
	}
}
