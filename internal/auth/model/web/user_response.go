package web

type UserRegisterRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"max=200,min=1" json:"password"`
	OpenId   string `validate:"required" json:"open_id"`
}
