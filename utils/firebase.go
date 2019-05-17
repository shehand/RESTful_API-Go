package utils

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
)

var client *db.Client

type Config struct {
	AuthOverride     *map[string]interface{} `json:"databaseAuthVariableOverride"`
	DatabaseURL      string                  `json:"databaseURL"`
	ProjectID        string                  `json:"projectId"`
	ServiceAccountID string                  `json:"serviceAccountId"`
	StorageBucket    string                  `json:"storageBucket"`
}

type App struct {}

func init(){
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL:"https://<CHANGE_ME>.firebaseio.com/",
	}

	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalf("firebase new app : %v", err)
	}

	client, err = app.Database(ctx)

	if err !=  nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

func NewApp(ctx context.Context, config *Config, opts ...option.ClientOption) (*App, error){}