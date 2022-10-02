package entity

import "github.com/google/uuid"

//ID entity
type ID uuid.UUID

//newID create a new entity ID
func NewID() ID{
	return ID(uuid.New())
}

func StringToID(s string) (ID, error){
	id,err := uuid.Parse(s)
	return ID(id),err
}