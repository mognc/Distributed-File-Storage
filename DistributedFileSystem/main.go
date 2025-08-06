package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	r := mux.NewRouter()

	// Basic routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/api/status", statusHandler).Methods("GET")

	// File system routes (placeholder for your distributed file system)
	r.HandleFunc("/api/files", filesHandler).Methods("GET")
	r.HandleFunc("/api/files/{filename}", fileHandler).Methods("GET", "POST", "DELETE")

	log.Printf("Distributed File System starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Distributed File System API",
		Status:  "running",
	}
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "System is healthy",
		Status:  "operational",
	}
	json.NewEncoder(w).Encode(response)
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Files endpoint - coming soon",
		Status:  "development",
	}
	json.NewEncoder(w).Encode(response)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	response := Response{
		Message: fmt.Sprintf("File operation for: %s", filename),
		Status:  "development",
	}
	json.NewEncoder(w).Encode(response)
}
