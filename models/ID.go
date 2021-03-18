package models

import "github.com/google/uuid"

//ID entity ID
type ID string

//create New entity ID
func NewID() ID {
	return ID(uuid.NewString())
}

//convert string to an entity ID
//func StringToID(s string) (ID, error){
//	id, err := uuid.Parse(s)
//	return  ID(id), err
//}
