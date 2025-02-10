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

// func (cfg *apiConfig) validateChirpHandler(w http.ResponseWriter, r *http.Request) {
// 	var input struct {
// 		Body string `json="body"`
// 	}
// 	var er struct{
// 		Error  string `json="error"`
// 	}
// 	var success struct {
// 		Valid bool
// 	}	
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err!=nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Header().Set("Content-type","application/json")
// 		er.Error = "Something went wrong"
// 		dat,err := json.Marshal(er)
// 		if err != nil {
// 			log.Printf("Error marshalling JSON: %s", err)
// 			w.WriteHeader(500)
// 			return
// 	}
// 		w.Write(dat)
// 		return
// 	}
// 	if len(input.Body) >40 {
// 			w.WriteHeader(400)
// 		w.Header().Set("Content-type","application/json")
// 		er.Error = "Chirp is too long"
// 		dat,err := json.Marshal(er)
// 		if err != nil {
// 			log.Printf("Error marshalling JSON: %s", err)
// 			w.WriteHeader(500)
// 			return
// 	}
// 		w.Write(dat)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-type","application/json")
// 		success.Valid = true
// 		dat,err := json.Marshal(success)
// 		if err != nil {
// 			log.Printf("Error marshalling JSON: %s", err)
// 			w.WriteHeader(500)
// 			return
// 	}
// 		w.Write(dat)
// }

func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{}
	mux := http.NewServeMux()

	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /api/validate_chirp", apiCfg.validateChirpHandler)

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