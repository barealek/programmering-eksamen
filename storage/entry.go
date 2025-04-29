package storage

type Entry struct {
	ID          string `json:"id"`
	Filnavn     string `json:"filnavn"`
	Sti         string `json:"sti"`
	Krypteret   bool   `json:"krypteret"`
	AdminSecret string `json:"admin_secret"`
}
