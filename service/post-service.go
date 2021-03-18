package service

import (
	m "Coursekita-chi/models"
	"github.com/pkg/errors"
	"Coursekita-chi/repository"
)

type IPostService interface {
	Validate(post *m.Post) error
	Create(post *m.Post) (*m.Post, error)
	FindAll()([]m.Post, error)
}

type service struct{}

var (
	repo repository.IPostRepo
)

func NewPostService(reposit repository.IPostRepo) IPostService {
	repo = reposit
	return &service{}
}

//method Validate
func (*service) Validate(post *m.Post) error{
	if post == nil {
		err := errors.New("The Post is empty")
		return err
	}

	if post.Title == "" {
 		err := errors.New("the post title is empty")
 		return err
	}

	return nil

}

//method FindALl
func (*service) Create(post *m.Post) (*m.Post, error){
	post.ID = m.NewID()
	return repo.Save(post)

}


//method Create
func (*service) FindAll()([]m.Post, error){
 	return repo.FindAll()
}



