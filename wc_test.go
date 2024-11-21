//what counts as whitespace characterr?
// test with spaces and tabs

package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestCountBytes(t *testing.T) {
	buffer := bytes.Buffer{}
	reader := bytes.NewReader(make([]byte, 10))
	readerSize := reader.Size()
	CountBytes(reader, &buffer)
	buffer_val, err := strconv.Atoi(buffer.String())
	if err != nil {
		t.Error("Error converting buffer value to string", err)
	}
	if buffer_val != int(readerSize) {
		t.Errorf("Expected byte count %d, got count %d", readerSize, buffer_val)
	}
}
