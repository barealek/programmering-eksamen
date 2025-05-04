package test

import (
	"net/http"
	"os"

	"github.com/barealek/programmering-eksamen/api"
	"github.com/barealek/programmering-eksamen/storage"
)

var h http.Handler

func init() {
	os.RemoveAll("data")
	os.Mkdir("data", 0755)

	st := storage.NewFsStorage("data/test_index.json")
	a := api.NewApi(st)
	h = a.Register()
}
