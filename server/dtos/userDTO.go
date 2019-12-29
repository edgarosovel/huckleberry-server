package dtos

type UserDTO struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
