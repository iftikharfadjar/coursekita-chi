package service

import (
	m "Coursekita-chi/models"
	"Coursekita-chi/repository"
	"github.com/pkg/errors"
)

type IPostService interface {
	Validate(post *m.Post) error
	Create(post *m.Post) (error)
	FindAll()([]m.Post, error)
	Remove(posts m.IDCollection) error
	Update(post *m.Post) error
	FindByID(post string) (*m.Post, error)
}

type service struct{}

var (
	repo repository.CrudInterface
)

func NewPostService(reposit repository.CrudInterface) IPostService {
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
func (*service) Create(post *m.Post) (error){
	post.ID = m.NewID()
	return repo.Save(post)

}

//method Create
func (*service) FindAll()([]m.Post, error){
	c, err := repo.FindAll("post")
 	posts := c.([]m.Post)
	return posts, err
}

func (*service) Remove(posts m.IDCollection ) error {
	err := repo.DeleteByID(posts.ID)
	return err
}

func (*service) Update(post *m.Post ) error {
	err := repo.UpdateByID(post)
	return err
}

func (*service) FindByID(postID string) (*m.Post, error){
	result , err := repo.FindByID(postID)
	post := result.(*m.Post)
	return post , err
}






