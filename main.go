package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Response struct {
	Service   string `json:"service"`
	RemoteURI string `json:"remoteURL"`
}

func handle(service string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		j := Response{
			Service:   service,
			RemoteURI: r.RequestURI,
		}

		json.NewEncoder(w).Encode(j)
	}
}

func main() {
	listenAddress := os.Getenv("ADDR")
	if listenAddress == "" {
		listenAddress = ":80"
	}

	service := os.Getenv("SERVICE")
	if service == "" {
		service = "n/a"
	}

	http.HandleFunc("/", handle(service))

	http.ListenAndServe(listenAddress, nil)
}
