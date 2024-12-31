package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func AccountRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /healthz", sendHealthz)

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
