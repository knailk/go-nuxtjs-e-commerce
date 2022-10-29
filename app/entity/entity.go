package entity

import (
	"strconv"

	"github.com/google/uuid"
)

//ID entity.
type ID uint32 

//newID create a new entity ID
func NewID() ID{
	id := uuid.New()
	return ID(id.ID())
}

func StringToID(s string) (ID, error){
	//id,err := uuid.Parse(s)
	id, err := strconv.ParseUint(s, 10, 64)
	return ID(id),err
}