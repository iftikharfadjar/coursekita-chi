package repository

import (
	"Coursekita-chi/models"
)

type IPostRepo interface {
	Save(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}
