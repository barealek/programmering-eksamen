package storage

import (
	"fmt"
)

type Entry struct {
	ID               string `json:"id"`
	Filnavn          string `json:"filnavn"`
	Krypteret        bool   `json:"krypteret"`
	AdminSecret      string `json:"admin_secret"`
	DownloadsTilbage int    `json:"downloads_tilbage"`
}

func (e *Entry) String() string {
	return fmt.Sprintf("ID:%s Filnavn:%s Krypteret:%t AdminSecret:%s DownloadsTilbage:%d", e.ID, e.Filnavn, e.Krypteret, e.AdminSecret, e.DownloadsTilbage)
}
