package utility

import (
	"context"

	"github.com/barealek/minilink/internal/database"
	"github.com/barealek/minilink/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllMinis() ([]models.Mini, error) {
	var minis []models.Mini

	coll := database.GetMongo().Collection("minis")

	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	cur.All(nil, &minis)

	return minis, nil

}

func GetMini(id uint64) models.Mini {
	var mini models.Mini

	coll := database.GetMongo().Collection("minis")

	mr := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	if mr.Err() != nil {
		panic(mr.Err())
	}
	mr.Decode(&mini)

	return mini
}
