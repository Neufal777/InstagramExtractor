package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TikTokDownloader/src"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	fh := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fh))

	r.HandleFunc("/", src.DownloadPage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("PORT=> ", port)
	http.ListenAndServe(":"+port, r)
}
