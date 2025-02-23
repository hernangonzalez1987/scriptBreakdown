package uuidgenerator

import "github.com/google/uuid"

type UUIDGenerator struct{}

func New() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (ref *UUIDGenerator) New() uuid.UUID {
	return uuid.New()
}
