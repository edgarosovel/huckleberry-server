package dtos

type NotificationDTO struct {
	ID             uint64            `json:"id"`
	BookmarksShare BookmarksShareDTO `json:"bookmark_share"`
}
