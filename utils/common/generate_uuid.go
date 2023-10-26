package common

import "github.com/google/uuid"

func GenerateID() uuid.UUID {
	return uuid.New()
}
