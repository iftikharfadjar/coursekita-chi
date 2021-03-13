package models

import "github.com/google/uuid"

//ID entity ID
type ID uuid.UUID

//create New entity ID
func NewID() ID{
	return ID(uuid.New())
}

//convert string to an entity ID
func StringToID(s string) (ID, error){
	id, err := uuid.Parse(s)
	return  ID(id), err
}

