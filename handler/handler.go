package handler

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
)

//go:embed docs/*
var docsFS embed.FS

func AccountRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /healthz", sendHealthz)

	// serve embedded api docs
	docsSubFS, err := fs.Sub(docsFS, "docs")
	if err != nil {
		log.Fatalf("failed to create sub filesystem: %v", err)
	}
	fileServer := http.FileServer(http.FS(docsSubFS))
	router.Handle("/docs/", http.StripPrefix("/docs", fileServer))

	return router
}

func sendHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)
	response["status"] = "ok"
	response["service"] = "account"

	res, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(res)
}
