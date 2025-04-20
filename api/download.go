package api

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (a *Api) Download(w http.ResponseWriter, r *http.Request) {
	navn := r.PathValue("filnavn")

	var readChain io.Reader

	// Disk
	fileSrc, err := os.Open("entries/" + navn)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer fileSrc.Close()
	readChain = fileSrc

	// Komprimering
	gzipReader, err := gzip.NewReader(readChain)
	if err != nil {
		http.Error(w, "Failed to create gzip reader", http.StatusInternalServerError)
		return
	}
	defer gzipReader.Close()
	readChain = gzipReader

	// Skriv
	n, err := io.Copy(w, readChain)

	if err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
		return
	}

	fmt.Println("Læste", n, "bytes fra disk til", r.RemoteAddr)
}
