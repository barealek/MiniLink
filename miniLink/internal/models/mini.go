package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mini struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`  // ID of the miniLink
	Redirect string             `json:"link" bson:"link"`         // Link to redirect to
	MiniLink string             `json:"minilink" bson:"minilink"` // The path of the minilink
	Clicks   uint64             `json:"clicks" bson:"clicks"`     // Amount of clicks
}
