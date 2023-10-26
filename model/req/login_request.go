package req

type LoginRequest struct {
	Identifier loginIdentifier
	Password   string
}

type loginIdentifier struct {
	Email string `json:"email" validate:"omitempty,email"`
	Name  string `json:"name" validate:"omitempty"`
}
