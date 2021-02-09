package initializemodels

import (
	"github.com/ceosss/lesson-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewLesson - Initializes a new lesson with the name provided
func NewLesson(name string) models.Lesson {
	lesson := models.Lesson{
		Name:     name,
		Labels:   []models.Label{},
		Models:   []primitive.ObjectID{},
		Question: []models.Question{},
	}
	return lesson
}
