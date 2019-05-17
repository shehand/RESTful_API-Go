package utils

import (
	"context"
	"log"
	"time"
	 firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

var client *db.Client

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