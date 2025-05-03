package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/barealek/programmering-eksamen/api"
	"github.com/barealek/programmering-eksamen/storage"
)

func main() {

	os.Mkdir("data", 0755)

	st := storage.NewFsStorage("data/dataindex.json")
	err := st.Load()
	if err != nil {
		fmt.Println("Failed to load storage", err)
		return
	}

	api := api.NewApi(st)

	http.ListenAndServe(":3000", api.Register())
}
