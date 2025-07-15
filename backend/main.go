package main

import (
	"log"
	"net/http"

	"womenNexus/firebase"
	"womenNexus/handlers"

	"github.com/gorilla/mux"
)

func main() {
	firebase.InitFirebase()
	defer firebase.FirestoreClient.Close()

	r := mux.NewRouter()

	// User Auth Routes
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/user/{uid}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user/{uid}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{uid}", handlers.DeleteUser).Methods("DELETE")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(r)))
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != "" {
			// Set CORS headers only when Origin is present (i.e., browser requests)
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Respond to preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
