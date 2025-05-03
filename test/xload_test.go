package test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/barealek/programmering-eksamen/api"
	"github.com/barealek/programmering-eksamen/storage"
)

// Denne pakke tester både up- og download-funktionen af filer

var h http.Handler

func TestDownload(t *testing.T) {
	mem := storage.NewMemStorage()
	a := api.NewApi(mem)
	h = a.Register()

	msg := []byte("dette er en testpakke")

	r, err := http.NewRequest("PUT", "/api/test.txt", bytes.NewBuffer(msg))
	if err != nil {
		t.Fatal(err)
	}

	w := newTestWriter()
	h.ServeHTTP(w, r)

	if w.status != 200 {
		t.Errorf("Expected status 200, got %d", w.status)
	}

	if !bytes.Equal(w.buf, msg) {
		t.Errorf("Expected %s, got %s", msg, w.buf)
	}
}
