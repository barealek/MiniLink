package utility

import (
	"context"

	"github.com/barealek/minilink/internal/database"
	"github.com/barealek/minilink/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetMini(miniId string) (models.Mini, error) {
	var mini models.Mini

	coll := database.GetMongo().Collection("minis")

	mr := coll.FindOne(context.TODO(), bson.D{{Key: "minilink", Value: miniId}})
	if mr.Err() != nil {
		return models.Mini{}, mr.Err()
	}
	mr.Decode(&mini)

	return mini, nil
}

func CreateMini(url string) (models.Mini, error) {
	mini := models.Mini{
		Redirect: url,
		MiniLink: GetRandomID(),
	}

	coll := database.GetMongo().Collection("minis")

	elem, err := coll.InsertOne(context.TODO(), mini)
	if err != nil {
		return models.Mini{}, err
	}

	mini.ID = elem.InsertedID.(primitive.ObjectID)

	return mini, nil
}

func IncrementClicks(miniLink string) error {
	coll := database.GetMongo().Collection("minis")

	_, err := coll.UpdateOne(context.TODO(), bson.D{{Key: "minilink", Value: miniLink}}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "clicks", Value: 1},
		}},
	})
	if err != nil {
		return err
	}

	return nil
}
