package api

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/barealek/programmering-eksamen/encryption"
)

func (a *Api) Download(w http.ResponseWriter, r *http.Request) {
	navn := r.PathValue("filnavn")
	krypteringsKey := r.URL.Query().Get("key")

	var readChain io.Reader

	// Disk
	fileSrc, err := os.Open("data/" + navn)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer fileSrc.Close()
	readChain = fileSrc

	// Kryptering
	if krypteringsKey != "" {

		fmt.Println("Dekryptering med key:", krypteringsKey)

		streamReader, err := encryption.EncryptedReader(krypteringsKey, readChain)
		if err != nil {
			http.Error(w, "Failed to create encrypted reader", http.StatusInternalServerError)
			return
		}
		readChain = streamReader
	} else {
		fmt.Println("Ingen dekryptering")
	}

	// Komprimering
	gzipReader, err := gzip.NewReader(readChain)
	if err != nil {
		// Der kommer ikke nogen fejl af at have en forkert dekrypteringskode
		// ovenfor. Hvis der er en fejl hernede, er det højst sandsynligt på
		// grund af, at der er en forkert dekrypteringskode ovenfor.

		http.Error(w, "Wrong decryption password", http.StatusForbidden)
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
