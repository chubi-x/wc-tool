//what counts as whitespace characterr?
// test with spaces and tabs

package main

import (
	"bytes"
	"testing"
)

func TestCountBytes(t *testing.T) {

	buf := bytes.NewBuffer(make([]byte, 5))

	count := Counter(buf, "bytes")
	bufLen := buf.Len()
	if count != bufLen {
		t.Errorf("Expected byte count %d, got count %d", bufLen, count)
	}
}
func TestCountLines(t *testing.T) {

	buf := bytes.NewBuffer([]byte("This is one line \n this is another line \n this is a third line"))
	lineCount := 3
	count := Counter(buf, "lines")

	if count != int(lineCount) {
		t.Errorf("Expected line count %d, got count %d", lineCount, count)
	}
}
func TestCountWords(t *testing.T) {
	buf := bytes.NewBuffer([]byte("There are four words"))
	wordCount := 4
	count := Counter(buf, "words")

	if count != int(wordCount) {
		t.Errorf("Expected word count %d, got count %d", wordCount, count)
	}
}
