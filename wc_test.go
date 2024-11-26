//what counts as whitespace characterr?
// test with spaces and tabs

package main

import (
	"bytes"
	"testing"
)

func TestCountBytes(t *testing.T) {

	buf := bytes.NewBuffer(make([]byte, 5))

	count := ByteCounter(buf)
	bufLen := buf.Len()
	if count != bufLen {
		t.Errorf("Expected byte count %d, got count %d", bufLen, count)
	}
}
func TestCountLines(t *testing.T) {

	buf := bytes.NewBuffer([]byte("This is one line \n this is another line \n this is a third line"))
	lineCount := 3
	count := LineCounter(buf)

	if count != int(lineCount) {
		t.Errorf("Expected line count %d, got count %d", lineCount, count)
	}
}
func TestCountWords(t *testing.T) {
	buf := bytes.NewBuffer([]byte("There are four words"))
	wordCount := 4
	count := WordCounter(buf)

	if count != int(wordCount) {
		t.Errorf("Expected word count %d, got count %d", wordCount, count)
	}
}
