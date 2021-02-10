package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model Structure
type Model struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required,min=3"`
	Location    string             `json:"location" bson:"location" validate:"url"`
	Coordinates Coordinates        `json:"coordinates" bson:"coordinates"`
}

// Lesson Structure
type Lesson struct {
	ID       primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name" validator:"required,min=3"`
	Labels   []Label              `json:"labels" bson:"labels"`
	Models   []primitive.ObjectID `json:"models" bson:"models"`
	Question []Question           `json:"questions" bson:"questions"`
}

// Question Structure
type Question struct {
	Statement   string      `json:"statement" bson:"statement" validator:"min=3"`
	Options     []Option    `json:"options" bson:"options"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates"`
}

// Option Structure
type Option struct {
	Option    string `json:"option" bson:"option" validator:"min=3"`
	IsCorrect bool   `json:"isCorrect" bson:"isCorrect"`
}

// Label struct
type Label struct {
	Label       string      `json:"label" bson:"label" validator:"min=3"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates"`
}

// Coordinates struct
type Coordinates struct {
	X int `json:"x" bson:"x"`
	Y int `json:"y" bson:"y"`
	Z int `json:"z" bson:"z"`
}

// User struct
type User struct {
	Email    string `json:"email" bson:"email" validate:"required,email,min=3"`
	Password string `json:"password" bson:"password" validate:"required,min=6"`
}

// Claims struct
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
