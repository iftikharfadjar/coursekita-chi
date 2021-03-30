package repository

import (
	"Coursekita-chi/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type postMongo struct{}

//new mongodb repository
func NewMongoDBRepository() CrudInterface {
	return &postMongo{}
}

//save Post
func (*postMongo) Save(c interface{}) (error) {
	post := c.(*models.Post)
	_, err := mg.Db.Collection("post").InsertOne(context.TODO(),post)

	if err != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
		return err
	}

	return nil
}


//find ALL Post
func (*postMongo) FindAll(c string) (interface{}, error) {

	var posts []models.Post
	cursor, err := mg.Db.Collection(c).Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return nil, err
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var post models.Post
		_ = cursor.Decode(&post)
		posts = append(posts, post)
	}

	return posts, nil
}

//Delete Post
func (*postMongo) DeleteByID(id []string) error {
	d := bson.D{{"id", bson.D{{"$in" , id }}}}
	var _, err = mg.Db.Collection("post").DeleteMany(context.TODO(), d)
	if err != nil {
		log.Printf("Error while deleting Documents, Reason: %v\n", err)
		return err
	}
	return nil
}

func (*postMongo) UpdateByID(update interface{}) error{
	post := update.(*models.Post)

	filter := bson.D{
		{"id", post.ID},
	}

	result := mg.Db.Collection("post").FindOneAndUpdate(context.TODO(),filter,bson.D{{"$set", post}})

	if result.Err() != nil {
		log.Printf("Error while updating Documents, Reason: %v\n", result.Err())
		return result.Err()
	}

	return  nil
}

func (*postMongo)  FindByID(c string) (interface{}, error) {
	var post models.Post

	err := mg.Db.Collection("post").FindOne(context.TODO(),bson.D{{"id",c},}).Decode(&post)

	if err !=  nil{
		log.Print("Error while updating Documents, Reason: %v\n", err)
		 return nil, err
	}

	return &post , nil
}
