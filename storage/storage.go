package storage

import "io"

type Storage interface {
	// Load indlæser data fra storage - typisk noget på disk som er permanent
	Load() error

	// Save gemmer data til storage - typisk noget på disk som er permanent
	Save() error

	Insert(*Entry) error
	Get(string) *Entry
	Delete(*Entry) error

	// FileDest returnerer en writer til en ny fil, der skal uploades
	FileDest(string) (io.WriteCloser, error)

	// FileSource returnerer en reader fra en eksisterende fil, der skal downloades
	FileSource(*Entry) (io.ReadCloser, error)
}
