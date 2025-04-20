package api

import (
	"net/http"

	"github.com/barealek/programmering-eksamen/storage"
)

type Api struct {
	st *storage.Storage
}

func NewApi(st *storage.Storage) *Api {
	return &Api{
		st: st,
	}
}

func (a *Api) Register() http.Handler {
	m := http.NewServeMux()

	m.HandleFunc("POST /api/upload/{filnavn}", a.Upload)
	m.HandleFunc("GET /api/download/{filnavn}", a.Download)

	return m
}
