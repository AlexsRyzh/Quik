package request

type RegisterRequest struct {
	Tag      string `json:"tag" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}
