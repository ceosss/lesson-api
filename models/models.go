package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Model Structure
type Model struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Location    string             `json:"location" bson:"location"`
	Coordinates []Coordinates      `json:"coordinates" bson:"coordinates"`
}

// Lesson Structure
type Lesson struct {
	ID       primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Labels   []string             `json:"labels" bson:"labels"`
	Models   []primitive.ObjectID `json:"models" bson:"models"`
	Question []Question           `json:"questions" bson:"questions"`
}

// Question Structure
type Question struct {
	Statement   string        `json:"statement" bson:"statement"`
	Options     []Option      `json:"options" bson:"options"`
	Coordinates []Coordinates `json:"coordinates" bson:"coordinates"`
}

// Option Structure
type Option struct {
	Option    string `json:"option" bson:"option"`
	IsCorrect bool   `json:"isCorrect" bson:"isCorrect"`
}

// Coordinates struct
type Coordinates struct {
	X int `json:"x" bson:"x"`
	Y int `json:"y" bson:"y"`
	Z int `json:"z" bson:"z"`
}
