package main

import (
	"net/http"
	"os"

	"github.com/barealek/programmering-eksamen/api"
	"github.com/barealek/programmering-eksamen/storage"
)

func main() {

	os.Mkdir("data", 0755)

	st := storage.NewStorage("data/dataindex.json")

	api := api.NewApi(st)

	http.ListenAndServe(":3000", api.Register())
}
