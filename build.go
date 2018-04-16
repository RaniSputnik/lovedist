package main

import (
	"fmt"
	"net/http"
)

func buildHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
