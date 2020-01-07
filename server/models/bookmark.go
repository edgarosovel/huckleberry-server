package models

import (
	"errors"

	"huckleberry.app/server/database"
	"huckleberry.app/server/dtos"
)

// Bookmark class to represent a URL bookmaRK
type Bookmark struct {
	BaseModel
	// User        User
	UserID      uint64 `sql:"type:bigint REFERENCES users(id)"`
	URL         string `json:"url"`
	ImgURL      string `json:"img_url"`
	Price       uint64 `json:"price"`
	Address     string `json:"address"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ToDTO returns a DTO representation
func (b *Bookmark) ToDTO() dtos.BookmarkDTO {
	bookmarkDTO := dtos.BookmarkDTO{
		ID:          b.ID,
		URL:         b.URL,
		ImgURL:      b.ImgURL,
		Price:       b.Price,
		Address:     b.Address,
		Title:       b.Title,
		Description: b.Description,
		CreatedAt:   b.CreatedAt,
	}

	return bookmarkDTO
}

func (b *Bookmark) Create() (*Bookmark, error) {
	err := database.DB.Debug().Create(&b).Error
	if err != nil {
		return &Bookmark{}, err
	}
	return b, nil
}

func (b *Bookmark) Delete() (*Bookmark, error) {
	if b.ID == 0 {
		return &Bookmark{}, errors.New("No ID provided")
	}
	err := database.DB.Debug().Delete(&b).Error
	if err != nil {
		return &Bookmark{}, err
	}
	return b, nil
}
