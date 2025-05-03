package api

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"

	"github.com/barealek/programmering-eksamen/encryption"
)

func (a *Api) Download(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	krypteringsKey := r.URL.Query().Get("key")
	fmt.Printf("id: %v\n", id)

	// Skaf entry file
	e := a.st.Get(id)
	if e == nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Slet entry file hvis der er 0 downloads tilbage
	if e.DownloadsTilbage == 0 {
		err := a.st.Delete(e)
		a.st.Save()
		if err != nil {
			fmt.Println("Failed to delete file", err)
			http.Error(w, "Internal error occured", http.StatusInternalServerError)
			return
		}
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Der er downloads tilbage, sæt det tilbage
	if e.DownloadsTilbage > 0 {
		e.DownloadsTilbage--
		a.st.Save()
	}

	var readCh io.Reader

	// Disk
	fileSrc, err := a.st.FileSource(e)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer fileSrc.Close()
	readCh = fileSrc

	// Kryptering
	if krypteringsKey != "" {

		fmt.Println("Dekryptering med key:", krypteringsKey)

		streamReader, err := encryption.EncryptedReader(krypteringsKey, readCh)
		if err != nil {
			http.Error(w, "Failed to create encrypted reader", http.StatusInternalServerError)
			return
		}
		readCh = streamReader
	} else {
		fmt.Println("Ingen dekryptering")
	}

	// Komprimering
	gzipReader, err := gzip.NewReader(readCh)
	if err != nil {
		// Der kommer ikke nogen fejl af at have en forkert dekrypteringskode
		// ovenfor. Hvis der er en fejl hernede, er det højst sandsynligt på
		// grund af, at der er en forkert dekrypteringskode ovenfor.

		http.Error(w, "Wrong decryption password", http.StatusForbidden)
		return
	}
	defer gzipReader.Close()
	readCh = gzipReader

	// Skriv
	n, err := io.Copy(w, readCh)

	if err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
		return
	}

	fmt.Println("Læste", n, "bytes fra disk til", r.RemoteAddr)
}
