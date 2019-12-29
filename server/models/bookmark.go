package models

// Bookmark class to represent a URL bookmaRK
type Bookmark struct {
	BaseModel
	User        User
	UserID      uint64 `sql:"type:bigint REFERENCES users(id)"`
	URL         string `json:"url"`
	Price       uint64 `json:"price"`
	Address     string `json:"address"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// func (b *Bookmark) Prepare() {
// 	b.ID = 0
// 	b.Title = html.EscapeString(strings.TrimSpace(b.Title))
// 	b.Description = html.EscapeString(strings.TrimSpace(b.Description))
// 	b.CreatedAt = time.Now()
// 	b.UpdatedAt = time.Now()
// }

// func (b *Bookmark) Validate() error {

// 	if b.Title == "" {
// 		return errors.New("Required Title")
// 	}
// 	if b.Description == "" {
// 		return errors.New("Required Description")
// 	}
// 	if b.UserID < 1 {
// 		return errors.New("Required Author")
// 	}
// 	return nil
// }

// func (b *Bookmark) SaveBookmark(db *gorm.DB) (*Bookmark, error) {
// 	var err error
// 	err = db.Debug().Model(&Bookmark{}).Create(&b).Error
// 	if err != nil {
// 		return &Bookmark{}, err
// 	}
// 	if b.ID != 0 {
// 		err = db.Debug().Model(&User{}).Where("id = ?", b.UserID).Take(&b.UserID).Error
// 		if err != nil {
// 			return &Bookmark{}, err
// 		}
// 	}
// 	return b, nil
// }

// func (b *Bookmark) FindAllBookmarks(db *gorm.DB) (*[]Bookmark, error) {
// 	var err error
// 	posts := []Bookmark{}
// 	err = db.Debug().Model(&Bookmark{}).Limit(100).Find(&posts).Error
// 	if err != nil {
// 		return &[]Bookmark{}, err
// 	}
// 	if len(posts) > 0 {
// 		for i, _ := range posts {
// 			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].UserID).Take(&posts[i].UserID).Error
// 			if err != nil {
// 				return &[]Bookmark{}, err
// 			}
// 		}
// 	}
// 	return &posts, nil
// }

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
