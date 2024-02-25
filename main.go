package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", slog.String("remote_address", r.RemoteAddr))
		content, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("cannot read the request body", slog.Any("error_message", err))
			http.Error(w, "cannot read the request body", http.StatusInternalServerError)
			return
		}
		fmt.Println(string(content))
		fmt.Fprint(w, "Request is logged")
	})

	server := &http.Server{
		Addr:    ":9999",
		Handler: mux,
	}

	slog.Info("server is listening on :9999")
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("cannot start the server", slog.Any("error_message", err))
	}
}
