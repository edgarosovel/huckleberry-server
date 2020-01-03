package models

import (
	"huckleberry.app/server/database"
	"huckleberry.app/server/dtos"
)

type BookmarksShare struct {
	BaseModel
	User       User
	UserID     uint64 `sql:"type:bigint REFERENCES users(id)"`
	Receiver   User
	ReceiverID uint64 `sql:"type:bigint REFERENCES users(id)"`
	IsAccepted bool
}

func FindSharesToUser(user User) []dtos.UserDTO {
	var usersDTO []dtos.UserDTO
	var shares []BookmarksShare

	database.DB.Debug().Where("receiver_id = ? AND is_accepted = true", user.ID).Find(&shares)
	for _, share := range shares {
		share.User.ID = share.UserID
		database.DB.Debug().First(&share.User)
		share.User.FindOwnBookmarks()
		usersDTO = append(usersDTO, share.User.ToDTO())
	}

	return usersDTO
}
