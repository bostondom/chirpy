package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

// func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
// 		cfg.fileserverHits.Add(1)
// 		next.ServeHTTP(w,r)
// 	})
// }

func (cfg *apiConfig) middlewareMetricsReset(){
	cfg.fileserverHits.Store(0)
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{}
	mux := http.NewServeMux()

	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("GET /healthz", handlerReadiness)
	mux.HandleFunc("POST /reset", apiCfg.handlerReset)
	mux.HandleFunc("GET /metrics", apiCfg.handlerMetrics)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

// func handlerReadiness(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(http.StatusText(http.StatusOK)))
// }