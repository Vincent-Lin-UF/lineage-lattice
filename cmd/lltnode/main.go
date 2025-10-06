package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Vincent-Lin-UF/lineage-lattice/internal/llt"
)

func main() {
	store := llt.NewStore()

	mux := http.NewServeMux()
	mux.HandleFunc("/llt/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.TrimPrefix(r.URL.Path, "/llt/")
		if key == "" {
			http.Error(w, "Missing keys in path: /llt/{key}", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodPut:
			defer r.Body.Close()
			val, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read body:"+err.Error(), http.StatusInternalServerError)
				return
			}

			store.Put(key, val)
			w.WriteHeader(http.StatusNoContent)

		case http.MethodGet:
			val, ok := store.Get(key)
			if !ok {
				http.NotFound(w, r)
				return
			}

			w.Header().Set("Content-Type", "application/octet-stream")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(val)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("LLT Node listening on: :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
