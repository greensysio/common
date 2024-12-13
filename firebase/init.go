package firebase

import (
	"context"
	"fmt"
	"os"

	"github.com/greensysio/common/log"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

// Firebase app
var Firebase *firebase.App

// DBClient for FB app
var DBClient *db.Client
var FSClient *firestore.Client

func init() {
	log.InitLogger(false)
}

// InitFirebaseApp init firebase app
func InitFirebaseApp() {
	projectId := os.Getenv("FB_PROJECT_ID")
	if projectId == "" {
		projectId = "EXGO-API"
	}
	conf := &firebase.Config{
		DatabaseURL: os.Getenv("FB_REALTIME_DB"),
		ProjectID:   projectId,
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(fmt.Sprintf("./env/%s/firebase_token.json", os.Getenv("INFO_ENV")))

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Error("Error initializing app:", err)
		os.Exit(1)
	}
	Firebase = app
	log.Info("Firebase App initialized!")
}

// GetDBClient returns Database client for Firebase app
func GetDBClient() *db.Client {
	if DBClient == nil {
		DBClient, err := Firebase.Database(context.Background())
		if err != nil {
			log.Error("Error initializing database client: ", err)
		}
		return DBClient
	}
	return DBClient
}

// GetDBClient returns Database client for Firebase app
func GetFSClient() *firestore.Client {
	if FSClient == nil {
		FSClient, err := Firebase.Firestore(context.Background())
		if err != nil {
			log.Error("Error initializing database client: ", err)
		}
		return FSClient
	}
	return FSClient
}
