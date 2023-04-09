package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	//return uuid.New()
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	//return uuid.Parse(id)
	id, err := uuid.Parse(s)
	return ID(id), err
}