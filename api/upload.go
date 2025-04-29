package api

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/barealek/programmering-eksamen/encryption"
	"github.com/barealek/programmering-eksamen/storage"
	"github.com/google/uuid"
)

func (a *Api) Upload(w http.ResponseWriter, r *http.Request) {

	navn := r.PathValue("filnavn")
	id := uuid.New().String()

	krypteringsKey := r.URL.Query().Get("key")

	var writeChain io.Writer
	defer r.Body.Close()

	// Disk
	fileDst, err := os.Create("data/" + id + ".bin")
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
		ID:          uuid.New().String(),
		Filnavn:     navn,
		Sti:         fmt.Sprintf("data/%s", id+".bin"),
		Krypteret:   krypteringsKey != "",
		AdminSecret: generateRandomString(12),
	}

	err = a.st.Insert(&elem)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	go a.st.Save()

	fmt.Println("Skrev", n, "bytes fra", r.RemoteAddr, "til disk")

	json.NewEncoder(w).Encode(map[string]string{
		"url":  "http://localhost:4321/download/" + elem.ID,
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
