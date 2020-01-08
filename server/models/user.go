package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"huckleberry.app/server/database"
	"huckleberry.app/server/dtos"
)

type User struct {
	BaseModel
	Name      string `gorm:"size:255;not null;" json:"name"`
	LastName  string `gorm:"size:255;not null;" json:"last_name"`
	Username  string `gorm:"size:255;not null;unique" json:"username" sql:"index"` // sql index for better query performance
	Email     string `gorm:"size:320;not null;unique" json:"email" sql:"index"`    // size 320 is max length for an email. Sql index for better query performance
	Password  string `gorm:"size:60;not null;" json:"password"`
	Bookmarks []Bookmark
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func UsernameExists(username string) (bool, error) {
	user := User{}
	err := database.DB.Debug().Where("username = ?", strings.ToLower(username)).First(&user).Error

	if err != nil {
		return false, err
	}

	return true, err
}

func EmailExists(email string) (bool, error) {
	user := User{}
	err := database.DB.Debug().Where("email = ?", email).First(&user).Error

	if err != nil {
		return false, err
	}

	return true, err
}

func (u *User) ToLoginDTO() dtos.UserLoginDTO {
	userLoginDTO := dtos.UserLoginDTO{
		Name:         u.Name,
		LastName:     u.LastName,
		CompleteName: fmt.Sprintf("%s %s", u.Name, u.LastName),
		Username:     u.Username,
		Email:        u.Email,
	}

	return userLoginDTO
}

func (u *User) ToDTO() dtos.UserDTO {
	bookmarksDTO := []dtos.BookmarkDTO{}
	for _, bookmark := range u.Bookmarks {
		bookmarksDTO = append(bookmarksDTO, bookmark.ToDTO())
	}

	userDTO := dtos.UserDTO{
		Name:         u.Name,
		LastName:     u.LastName,
		CompleteName: fmt.Sprintf("%s %s", u.Name, u.LastName),
		Username:     u.Username,
		Bookmarks:    bookmarksDTO,
	}

	return userDTO
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) FindByUsername() {
	database.DB.Debug().Where("username = ?", strings.ToLower(u.Username)).First(&u)
}

// FindOwnBookmarks by ID or Username
func (u *User) FindOwnBookmarks() error {
	var err error
	// var bookmarks []Bookmark
	if u.ID == 0 && u.Username == "" {
		return errors.New("ID or Username not found")
	}
	if u.ID == 0 {
		err = database.DB.Debug().Where("username = ?", u.Username).First(&u).Error
	}
	if err != nil {
		return err
	}
	database.DB.Debug().Model(&u).Related(&u.Bookmarks)
	// Preload
	// u.Bookmarks = FindBookmarksByID(u.ID)
	return err
}

func (u *User) ValidateForCreate() error {
	if u.Name == "" {
		return errors.New("Required name")
	}
	if u.LastName == "" {
		return errors.New("Required last name")
	}
	if u.Username == "" {
		return errors.New("Required username")
	} else {
		u.Username = strings.ToLower(u.Username)
	}
	if u.Password == "" {
		return errors.New("Required password")
	}
	if u.Email == "" {
		return errors.New("Required email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid email")
	}
	return nil

}

// func (u *User) ValidateForUpdate() error {
// 	if u.Username == "" {
// 		return errors.New("Required Username")
// 	}
// 	if u.Password == "" {
// 		return errors.New("Required Password")
// 	}
// 	if u.Email == "" {
// 		return errors.New("Required Email")
// 	}
// 	if err := checkmail.ValidateFormat(u.Email); err != nil {
// 		return errors.New("Invalid Email")
// 	}

// 	return nil
// }

// func (u *User) ValidateForLogin() error {
// 	if u.Password == "" {
// 		return errors.New("Required Password")
// 	}
// 	if u.Email == "" {
// 		return errors.New("Required Email")
// 	}
// 	if err := checkmail.ValidateFormat(u.Email); err != nil {
// 		return errors.New("Invalid Email")
// 	}
// 	return nil
// }

// Function to save a User in the DB
func (u *User) Create() (*User, error) {
	var err error

	err = u.ValidateForCreate()
	if err != nil {
		return &User{}, err
	}

	err = database.DB.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// Function to list all users. For development purposes only
// func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
// 	var err error
// 	users := []User{}
// 	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
// 	if err != nil {
// 		return &[]User{}, err
// 	}
// 	return &users, err
// }

// Function to find a user by its ID
// func (u *User) FindUserByID(db *gorm.DB, uid uint64) (*User, error) {
// 	var err error
// 	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	if gorm.IsRecordNotFoundError(err) {
// 		return &User{}, errors.New("User Not Found")
// 	}
// 	return u, err
// }

// func (u *User) UpdateUserByID(db *gorm.DB, uid uint64) (*User, error) {

// 	// To hash the password
// 	err := u.BeforeSave()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
// 		map[string]interface{}{
// 			"password":  u.Password,
// 			"name":      u.Name,
// 			"last_name": u.LastName,
// 			"username":  u.Username,
// 			"email":     u.Email,
// 			"update_at": time.Now(),
// 		},
// 	)
// 	if db.Error != nil {
// 		return &User{}, db.Error
// 	}
// 	// This is the display the updated user
// 	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }

// func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

// 	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

// 	if db.Error != nil {
// 		return 0, db.Error
// 	}
// 	return db.RowsAffected, nil
// }
