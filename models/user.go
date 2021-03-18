package models

import "time"

//User Struct
type User struct {
	ID        ID
	Username  string
	email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts     []ID
}

//create New User

//Get User by ID

//Validate User

//Validate user password

//generate password

//add post depend by user

//remove post depend by user

//get post depend by user
