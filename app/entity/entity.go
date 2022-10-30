package entity

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

//ID entity.
type ID uint32 

//NewID create a new entity ID
func NewID() ID{
	id := uuid.New()
	return ID(id.ID())
}

//StringToID convert string to ID
func StringToID(s string) (ID, error){
	//id,err := uuid.Parse(s)
	id, err := strconv.ParseUint(s, 10, 64)
	return ID(id),err
}

//StringToTime convert string to time.time //2022-10-29 10:29:28.8512522+07:00 
func StringToTime(s string) (time.Time, error){
	return time.Parse(time.RFC3339, s)
}