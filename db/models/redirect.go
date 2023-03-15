package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Redirect struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	ActiveLink  string             `json:"active_link" bson:"active_link"`
	HistoryLink string             `json:"history_link" bson:"history_link"`
}
