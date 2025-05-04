package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// Denne pakke tester både up- og download-funktionen af filer

func TestUpDownload(t *testing.T) {

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

	var res = make(map[string]string)
	if err := json.NewDecoder(w.buf).Decode(&res); err != nil {
		t.Error("Failed to decode response", err)
	}

	var (
		// code string
		id string
		ok bool
	)
	if id, ok = res["id"]; !ok {
		t.Error("Expected id in response")
	}
	if _, ok = res["code"]; !ok {
		t.Error("Expected code in response")
	}

	fmt.Printf("id: %v\n", id)

	// Test download
	r, err = http.NewRequest("GET", "/api/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	w = newTestWriter()
	h.ServeHTTP(w, r)

	if w.status != 200 {
		t.Errorf("Expected status 200, got %d", w.status)
	}
}
