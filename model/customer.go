package model

type Customer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerData struct {
	Customers []*Customer `json:"customers"`
}
