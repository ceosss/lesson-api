package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Model Structure
type Model struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Location    string             `bson:"location,omitempty"`
	Coordinates []Coordinates      `bson:"coordinates,omitempty"`
}

// Coordinates struct
type Coordinates struct {
	X int `bson:"x,omitempty"`
	Y int `bson:"y,omitempty"`
	Z int `bson:"z,omitempty"`
}
