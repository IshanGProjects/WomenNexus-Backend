package firebase

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var App *firebase.App
var AuthClient *auth.Client
var FirestoreClient *firestore.Client

func InitFirebase() {
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	var err error
	App, err = firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}

	AuthClient, err = App.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing firebase auth: %v", err)
	}

	FirestoreClient, err = App.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing firestore: %v", err)
	}
}
