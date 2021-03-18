package repository

import (
	"Coursekita-chi/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"log"
)

type repo struct{}

//new mongodb repository
func NewMongoDBRepository() IPostRepo {
	return &repo{}
}

// DATABASE INSTANCE


func init(){
	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
}


//save Post
func (*repo) Save(post *models.Post) (*models.Post, error) {

	_, err := mg.Db.Collection("post").InsertOne(context.TODO(), *post)

	if err != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
		return nil, err
	}

	return post, nil
}

//find ALL Post
func (*repo) FindAll() ([]models.Post, error) {

	posts := []models.Post{}
	cursor, err := mg.Db.Collection("post").Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return nil, err
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	return posts, nil
}
