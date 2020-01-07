package models

import (
	"errors"

	"huckleberry.app/server/database"
	"huckleberry.app/server/dtos"
)

type BookmarksShare struct {
	BaseModel
	User       User   `gorm:"-" sql:"-"`
	UserID     uint64 `sql:"type:bigint REFERENCES users(id)"`
	Receiver   User   `gorm:"-" sql:"-" json:"receiver"`
	ReceiverID uint64 `sql:"type:bigint REFERENCES users(id)"`
	IsAccepted bool
}

func (bs *BookmarksShare) ToDTO() dtos.BookmarksShareDTO {
	bookmarksShareDTO := dtos.BookmarksShareDTO{
		ID:         bs.ID,
		User:       bs.User.ToDTO(),
		Receiver:   bs.Receiver.ToDTO(),
		IsAccepted: bs.IsAccepted,
	}

	return bookmarksShareDTO
}

func (bs *BookmarksShare) Create() (*BookmarksShare, error) {
	tx := database.DB.Debug().Begin()
	// Save share intent
	bs.IsAccepted = false
	if err := tx.Create(&bs).Error; err != nil {
		tx.Rollback()
		return &BookmarksShare{}, err
	}
	// Create a notification for recipient
	notification := Notification{
		UserID:           bs.ReceiverID,
		BookmarksShareID: bs.ID,
	}

	if err := tx.Create(&notification).Error; err != nil {
		tx.Rollback()
		return &BookmarksShare{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &BookmarksShare{}, err
	}

	return bs, nil
}

func (bs *BookmarksShare) FindByUserAndReceiver() bool {
	var count int
	database.DB.Model(&bs).Where("user_id = ? AND receiver_id = ?", bs.UserID, bs.ReceiverID).Count(&count)
	if count == 0 {
		return false
	}
	return true
}

func (bs *BookmarksShare) Delete() (*BookmarksShare, error) {
	if bs.ID == 0 {
		return &BookmarksShare{}, errors.New("No ID provided")
	}
	if err := database.DB.Delete(&bs).Error; err != nil {
		return &BookmarksShare{}, err
	}
	return bs, nil
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

func FindSharesFromUser(user User) []dtos.BookmarksShareDTO {
	var bookmarksShareDTO []dtos.BookmarksShareDTO
	var shares []BookmarksShare

	database.DB.Debug().Where("user_id = ?", user.ID).Find(&shares)
	for _, share := range shares {
		share.Receiver.ID = share.ReceiverID
		database.DB.Debug().First(&share.Receiver)
		bookmarksShareDTO = append(bookmarksShareDTO, share.ToDTO())
	}

	return bookmarksShareDTO
}
