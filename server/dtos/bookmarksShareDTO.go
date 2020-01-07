package dtos

type BookmarksShareDTO struct {
	ID         uint64  `json:"id"`
	User       UserDTO `json:"user"`
	Receiver   UserDTO `json:"receiver"`
	IsAccepted bool    `json:"is_accepted"`
}
