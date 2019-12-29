package models

type BookmarksShare struct {
	BaseModel
	User       User
	UserID     uint64 `sql:"type:bigint REFERENCES users(id)"`
	Receiver   User
	ReceiverID uint64 `sql:"type:bigint REFERENCES users(id)"`
	IsAccepted bool
}
