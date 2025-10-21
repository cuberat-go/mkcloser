package closer

import (
	// Built-in/core modules.
	"io"
)

// WrappedReader wraps an io.Reader to provide a no-op Close method.
type WrappedReader struct {
	reader io.Reader
}

// WrapReader wraps an io.Reader and returns an io.ReadCloser with a no-op Close
// method.
func WrapReader(r io.Reader) io.ReadCloser {
	return &WrappedReader{reader: r}
}

// Read reads up to len(p) bytes into p. See io.Reader.Read for more details.
func (wr *WrappedReader) Read(p []byte) (n int, err error) {
	return wr.reader.Read(p)
}

// Close is a no-op that satisfies the io.Closer interface.
func (wr *WrappedReader) Close() error {
	return nil
}

// WrappedWriter wraps an io.Writer to provide a no-op Close method.
type WrappedWriter struct {
	writer io.Writer
}

// WrapWriter wraps an io.Writer and returns an io.WriteCloser with a no-op
// Close method.
func WrapWriter(w io.Writer) io.WriteCloser {
	return &WrappedWriter{writer: w}
}

// Write writes len(p) bytes from p to the underlying data stream. See
// io.Writer.Write for more details.
func (ww *WrappedWriter) Write(p []byte) (n int, err error) {
	return ww.writer.Write(p)
}

// Close is a no-op that satisfies the io.Closer interface.
func (ww *WrappedWriter) Close() error {
	return nil
}
