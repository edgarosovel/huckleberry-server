package dtos

import "time"

type BookmarksResponseDTO struct {
	OwnBookmarks []BookmarkDTO `json:"my_bookmarks"`
	Shared       []UserDTO     `json:"shared"`
}

type BookmarkDTO struct {
	ID          uint64    `json:"id"`
	URL         string    `json:"url"`
	ImgURL      string    `json:"img_url"`
	Price       uint64    `json:"price"`
	Address     string    `json:"address"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
