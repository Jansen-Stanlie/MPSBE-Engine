package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Config struct {
	Server   *http.Server
	Database *mongo.Client
}
