package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const googleURL = "https://www.google.com"

var client = &http.Client{Timeout: 10 * time.Second}

type apiResponse struct {
	URL        string `json:"url,omitempty"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code,omitempty"`
	Error      string `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, statusCode int, response apiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func googleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeJSON(w, http.StatusMethodNotAllowed, apiResponse{
			Status: "error",
			Error:  "only GET requests are allowed",
		})
		return
	}

	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, googleURL, nil)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, apiResponse{Status: "error", Error: err.Error()})
		return
	}

	response, err := client.Do(req)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, apiResponse{
			URL: googleURL, Status: "error", Error: err.Error(),
		})
		return
	}
	defer response.Body.Close()

	writeJSON(w, http.StatusOK, apiResponse{
		URL: googleURL, Status: response.Status, StatusCode: response.StatusCode,
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, apiResponse{Status: "ok"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/google", googleHandler)
	mux.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("REST API listening on http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
