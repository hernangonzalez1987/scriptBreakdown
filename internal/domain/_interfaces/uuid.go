package _interfaces

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() uuid.UUID
}
