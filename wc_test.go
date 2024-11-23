//what counts as whitespace characterr?
// test with spaces and tabs

package main

import (
	"bytes"
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
