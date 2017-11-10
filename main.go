package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)

		}
		defer f.Close()
		var contentType string
		path := r.URL.Path
		switch {
		case strings.HasSuffix(path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(path, "png"):
			contentType = "image/png"
		default:
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)
	})

	http.ListenAndServe(":8000", nil)
}
