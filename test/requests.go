package test

import (
	"bytes"
	"net/http"
)

type TestWriter struct {
	buf    *bytes.Buffer
	header http.Header
	status int
}

func (w *TestWriter) Header() http.Header {
	return w.header
}

func (w *TestWriter) Write(b []byte) (int, error) {
	return w.buf.Write(b)
}

func (w *TestWriter) WriteHeader(status int) {
	w.status = status
}

func (w *TestWriter) Result() (*bytes.Buffer, int) {
	return w.buf, w.status
}

func newTestWriter() *TestWriter {
	return &TestWriter{
		buf:    new(bytes.Buffer),
		header: make(http.Header),
		status: 200,
	}
}
