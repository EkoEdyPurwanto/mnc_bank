package req

type RegisterRequest struct {
	Name            string `json:"name" validate:"required,min=3,max=30"`
	Email           string `json:"email" validate:"email"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,eqfield=Password"`
}
