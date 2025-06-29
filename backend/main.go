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
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	// User Auth Routes
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/user/{uid}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user/{uid}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{uid}", handlers.DeleteUser).Methods("DELETE")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
