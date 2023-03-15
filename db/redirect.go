package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (instance *Instance) IsActiveLinkExists(link string) (bool, error) {
	filter := bson.M{"active_link": link}
	count, _ := instance.Collection.CountDocuments(context.Background(), filter)
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (instance *Instance) IsHistoryLinkExists(link string) (bool, error) {
	filter := bson.M{"history_link": link}
	count, _ := instance.Collection.CountDocuments(context.Background(), filter)
	if count > 0 {
		return true, nil
	}
	return false, nil
}
