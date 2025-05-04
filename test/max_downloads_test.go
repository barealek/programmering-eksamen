package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// Denne pakke tester både up- og download-funktionen af filer

func TestMaxDownloads(t *testing.T) {

	msg := []byte("dette er en testpakke, som kun kan downloades 3 gange")
	var downloads = "3"

	r, err := http.NewRequest("PUT", "/api/krypteret.txt?downloads="+downloads, bytes.NewBuffer(msg))
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
	for i := range 4 {
		r, err = http.NewRequest("GET", "/api/"+id, nil)
		if err != nil {
			t.Fatal(err)
		}

		w = newTestWriter()
		h.ServeHTTP(w, r)

		if i == 3 {
			if w.status != http.StatusNotFound { // Efter 3 downloads (0, 1, 2, _3_) bør filen være slettet
				t.Errorf("Expected status 403, got %d", w.status)
			}
		} else {

			if w.status != 200 {
				t.Errorf("Expected status 200, got %d", w.status)
			}

			if bytes, err := io.ReadAll(w.buf); err != nil || string(bytes) != string(msg) {
				t.Errorf("Expected %s, got %s", string(msg), string(bytes))
			}
		}
	}

}
