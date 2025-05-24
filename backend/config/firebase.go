package config

import (
	"context"
	"log"
	"sync"
	"fmt"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var (
	FirebaseAuth *auth.Client
	initOnce     sync.Once
)
var (
	FirestoreClient *firestore.Client
	initFirestoreOnce sync.Once
)


func InitFirebase() error {
	initOnce.Do(func() { // ✅ Ensures FirebaseAuth initializes ONCE
		opt := option.WithCredentialsFile("E:/GoLang/ieltsAiTutor/backend/config/serviceAccountKey.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("Error initializing Firebase: %v", err)
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("Error initializing Firebase Auth: %v", err)
		}

		FirebaseAuth = client
		log.Println("✅ Firebase Auth initialized successfully!")
	})

	if FirebaseAuth == nil {
		return fmt.Errorf("FirebaseAuth is still nil after InitFirebase execution")
	}

	return nil
}


// firestore

func InitFirestore() error {
	initFirestoreOnce.Do(func() { // ✅ Ensures Firestore initializes ONCE per test session
		opt := option.WithCredentialsFile("E:/GoLang/ieltsAiTutor/backend/config/serviceAccountKey.json")
		client, err := firestore.NewClient(context.Background(), "ieltsaitutorfirebase", opt)
		if err != nil {
			log.Fatalf("Error initializing Firestore: %v", err)
		}
		FirestoreClient = client
		log.Println("✅ Firestore initialized successfully!")
	})

	if FirestoreClient == nil {
		return fmt.Errorf("FirestoreClient is still nil after InitFirestore execution")
	}

	return nil
}