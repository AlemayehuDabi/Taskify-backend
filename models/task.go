package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID primitive.ObjectID `bson: "_id,omitempty", json: "id"`
	Body string `bson: "body", json: "body"`
	Completed bool `bson: "completed", json: "completed"`

}