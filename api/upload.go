package api

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (a *Api) Upload(w http.ResponseWriter, r *http.Request) {

	navn := r.PathValue("filnavn")

	var writeChain io.Writer
	defer r.Body.Close()

	// Disk
	fileDst, err := os.Create("entries/" + navn)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer fileDst.Close()
	writeChain = fileDst

	// Komprimering
	gzipWriter := gzip.NewWriter(writeChain)
	defer gzipWriter.Close()
	writeChain = gzipWriter

	// Skriv
	n, err := io.Copy(writeChain, r.Body)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	fmt.Println("Skrev", n, "bytes fra", r.RemoteAddr, "til disk")
}
