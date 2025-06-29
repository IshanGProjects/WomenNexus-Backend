package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"womenNexus/firebase"
	"womenNexus/models"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req struct {
		Email    string             `json:"email"`
		Password string             `json:"password"`
		Profile  models.UserProfile `json:"profile"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	userRecord, err := firebase.AuthClient.CreateUser(ctx, (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password),
	)
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = firebase.FirestoreClient.Collection("users").Doc(userRecord.UID).Set(ctx, req.Profile)
	if err != nil {
		http.Error(w, "Failed to save profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"uid": userRecord.UID})
}

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	idToken := r.Header.Get("Authorization")
	if idToken == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	token, err := firebase.AuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"uid": token.UID})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	uid := mux.Vars(r)["uid"]

	doc, err := firebase.FirestoreClient.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(doc.Data())
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	uid := mux.Vars(r)["uid"]

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err := firebase.FirestoreClient.Collection("users").Doc(uid).Set(ctx, updates, firestore.MergeAll)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	uid := mux.Vars(r)["uid"]

	// Delete from Firebase Auth
	if err := firebase.AuthClient.DeleteUser(ctx, uid); err != nil {
		http.Error(w, "Failed to delete auth user", http.StatusInternalServerError)
		return
	}

	// Delete from Firestore
	if _, err := firebase.FirestoreClient.Collection("users").Doc(uid).Delete(ctx); err != nil {
		http.Error(w, "Failed to delete Firestore profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}
