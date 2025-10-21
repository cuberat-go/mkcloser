package mkcloser_test

import (
	"bytes"
	"io"
	"testing"

	closer "github.com/cuberat-go/mkcloser"
)

func TestWrapReader(t *testing.T) {
	originalData := []byte("Hello, World!")
	reader := bytes.NewReader(originalData)
	wrappedReader := closer.WrapReader(reader)

	readData := make([]byte, len(originalData))
	n, err := wrappedReader.Read(readData)
	if err != nil && err != io.EOF {
		t.Fatalf("Unexpected error during Read: %v", err)
	}
	if n != len(originalData) {
		t.Fatalf("Expected to read %d bytes, got %d bytes", len(originalData), n)
	}
	if !bytes.Equal(readData, originalData) {
		t.Fatalf("Read data does not match original data")
	}

	if err := wrappedReader.Close(); err != nil {
		t.Fatalf("Close should be a no-op, but got error: %v", err)
	}
}

func TestWrapWriter(t *testing.T) {
	originalData := []byte("Hello, Writer!")
	var buffer bytes.Buffer
	wrappedWriter := closer.WrapWriter(&buffer)

	n, err := wrappedWriter.Write(originalData)
	if err != nil {
		t.Fatalf("Unexpected error during Write: %v", err)
	}
	if n != len(originalData) {
		t.Fatalf("Expected to write %d bytes, wrote %d bytes", len(originalData), n)
	}
	if !bytes.Equal(buffer.Bytes(), originalData) {
		t.Fatalf("Written data does not match original data")
	}

	if err := wrappedWriter.Close(); err != nil {
		t.Fatalf("Close should be a no-op, but got error: %v", err)
	}
}
