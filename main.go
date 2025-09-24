package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type health struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(health{
			Status:  "OK",
			Version: version,
			Commit:  commit,
			Date:    date,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("gofafo %s (%s) listening on :%s\n", version, commit, port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}
