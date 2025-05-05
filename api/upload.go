package api

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/barealek/programmering-eksamen/encryption"
	"github.com/barealek/programmering-eksamen/storage"
	"github.com/google/uuid"
)

func (a *Api) Upload(w http.ResponseWriter, r *http.Request) {

	// Hent vigtige informationer fra HTTP-requesten
	navn := r.PathValue("filnavn")
	krypteringsKey := r.URL.Query().Get("key")
	antalDownloadsStr := r.URL.Query().Get("downloads")
	antalDownloads, err := strconv.Atoi(antalDownloadsStr) // Brug strconv til at parse antalDownloads fra en string til et heltal
	if antalDownloadsStr != "" && err != nil {
		http.Error(w, "Set downloads to a number", http.StatusBadRequest)
		return
	}

	if antalDownloadsStr == "" {
		antalDownloads = -1 // Hvis antalDownloads ikke er angivet, sæt det til -1, som vi har defineret som uendeligt downloads
	}

	// Lav en io.Writer
	var writeChain io.Writer
	defer r.Body.Close()

	// Disk
	id := uuid.New().String()
	fileDst, err := a.st.FileDest(id)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer fileDst.Close()
	writeChain = fileDst

	// Kryptering
	if krypteringsKey != "" {

		fmt.Println("Kryptering med key:", krypteringsKey)

		streamWriter, err := encryption.EncryptedWriter(krypteringsKey, writeChain)
		if err != nil {
			http.Error(w, "Failed to create encrypted writer", http.StatusInternalServerError)
			return
		}
		defer streamWriter.Close()
		writeChain = streamWriter

	} else {
		fmt.Println("Ingen kryptering")
	}

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

	elem := storage.Entry{
		ID:               id,
		Filnavn:          navn,
		Krypteret:        krypteringsKey != "",
		AdminSecret:      generateRandomString(12),
		DownloadsTilbage: antalDownloads,
	}

	err = a.st.Insert(&elem)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	a.st.Save()

	fmt.Println("Skrev", n, "bytes fra", r.RemoteAddr, "til disk")

	json.NewEncoder(w).Encode(map[string]string{
		"id":   elem.ID,
		"code": elem.AdminSecret,
	})
}

func generateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
