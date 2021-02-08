package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Model Structure
type Model struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Location    string             `json:"location" bson:"location,omitempty"`
	Coordinates []Coordinates      `json:"coordinates" bson:"coordinates,omitempty"`
}

// Coordinates struct
type Coordinates struct {
	X int `json:"x" bson:"x,omitempty"`
	Y int `json:"y" bson:"y,omitempty"`
	Z int `json:"z" bson:"z,omitempty"`
}
