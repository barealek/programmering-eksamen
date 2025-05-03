package api

import (
	"fmt"
	"net/http"
)

func (a *Api) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	e := a.st.Get(id)
	if e == nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	clientSecret := r.URL.Query().Get("secret")
	if clientSecret != e.AdminSecret {
		http.Error(w, "Wrong secret", http.StatusForbidden)
		return
	}

	err := a.st.Delete(e)
	if err != nil {
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		return
	}

	err = a.st.Save()
	if err != nil {
		fmt.Println("Failed to save storage", err)
		http.Error(w, "An internal error occured", http.StatusInternalServerError)
		return
	}

	fmt.Println("Slettede", id, "fra", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
