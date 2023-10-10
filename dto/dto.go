package dto

type DTOUsers struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}