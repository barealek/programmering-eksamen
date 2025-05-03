package test

import "net/http"

type TestWriter struct {
	buf    []byte
	header http.Header
	status int
}

func (w *TestWriter) Header() http.Header {
	return w.header
}

func (w *TestWriter) Write(b []byte) (int, error) {
	w.buf = append(w.buf, b...)
	return len(b), nil
}

func (w *TestWriter) WriteHeader(status int) {
	w.status = status
}

func (w *TestWriter) Result() ([]byte, int) {
	return w.buf, w.status
}

func newTestWriter() *TestWriter {
	return &TestWriter{
		buf:    make([]byte, 0),
		header: make(http.Header),
		status: 0,
	}
}
