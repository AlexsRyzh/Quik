package request

type UserRequest struct {
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	ImgLink string `json:"login"`
}
