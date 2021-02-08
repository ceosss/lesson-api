package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Model Structure
type Model struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Location    string             `bson:"location,omitempty"`
	Coordinates []int              `bson:"coordinated,omitempty"`
}
