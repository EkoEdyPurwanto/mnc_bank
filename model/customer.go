package model

import "github.com/google/uuid"

type Customer struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type CustomerData struct {
	Customers []*Customer `json:"customers"`
}
