package models

import (
	"errors"

	"huckleberry.app/server/database"
	"huckleberry.app/server/dtos"
)

type Notification struct {
	BaseModel
	User             User
	UserID           uint64 `sql:"type:bigint REFERENCES users(id)"`
	BookmarksShare   BookmarksShare
	BookmarksShareID uint64 `sql:"type:bigint REFERENCES bookmarks_shares(id)"`
}

func (n *Notification) DeleteByBookmarksShareID() error {
	err := database.DB.Where("bookmarks_share_id = ?", n.BookmarksShareID).Delete(&Notification{}).Error
	return err
}

func (n *Notification) AcceptBookmarksShare() error {
	if n.ID == 0 {
		return errors.New("No ID specified")
	}

	tx := database.DB.Debug().Begin()

	// get notification
	if err := tx.First(&n).Error; err != nil {
		return err
	}

	// get its bookmark share
	if err := tx.Model(&n).Related(&n.BookmarksShare).Error; err != nil {
		return err
	}

	// Patch share
	n.BookmarksShare.IsAccepted = true
	if err := tx.Save(&n.BookmarksShare).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete notification
	if err := tx.Delete(&n).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (n *Notification) RejectBookmarksShare() error {
	if n.ID == 0 {
		return errors.New("No ID specified")
	}

	tx := database.DB.Debug().Begin()

	// get notification
	if err := tx.First(&n).Error; err != nil {
		return err
	}

	// get its bookmark share
	if err := tx.Model(&n).Related(&n.BookmarksShare).Error; err != nil {
		return err
	}

	// Delete notification
	if err := tx.Delete(&n).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete share
	n.BookmarksShare.IsAccepted = true
	if err := tx.Delete(&n.BookmarksShare).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (n *Notification) ToDTO() dtos.NotificationDTO {
	notificationDTO := dtos.NotificationDTO{
		ID:             n.ID,
		BookmarksShare: n.BookmarksShare.ToDTO(),
	}

	return notificationDTO
}

func FindNotificationsByUser(u *User) ([]dtos.NotificationDTO, error) {
	var notifications []Notification
	err := database.DB.Debug().Where("user_id = ?", u.ID).Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	var notificationsDTO []dtos.NotificationDTO
	for _, notification := range notifications {
		database.DB.Debug().Model(&notification).Related(&notification.BookmarksShare)
		database.DB.Debug().Model(&notification.BookmarksShare).Related(&notification.BookmarksShare.User)
		notificationsDTO = append(notificationsDTO, notification.ToDTO())
	}

	return notificationsDTO, err
}
