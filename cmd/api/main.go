package main

import (
	"net/http"

	"github.com/barealek/programmering-eksamen/api"
	"github.com/barealek/programmering-eksamen/storage"
)

func main() {

	st := storage.NewStorage("dataindex.json")

	api := api.NewApi(st)

	http.ListenAndServe(":3000", api.Register())
}
