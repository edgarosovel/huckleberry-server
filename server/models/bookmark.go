package models

import (
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

// ToDTO Function to return a DTO representation
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

// func (b *Bookmark) DeleteBookmark(db *gorm.DB, ID uint) (int, error) {

// 	db = db.Debug().Delete(&Bookmark{ID: ID})

// 	if db.Error != nil {
// 		if gorm.IsRecordNotFoundError(db.Error) {
// 			return 0, errors.New("Bookmark not found")
// 		}
// 		return 0, db.Error
// 	}
// 	return db.RowsAffected, nil
// }
