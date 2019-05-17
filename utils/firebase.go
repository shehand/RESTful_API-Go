package utils

import (
	"context"
	"log"
	"time"
	 firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

var client *db.Client

func init()