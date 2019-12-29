package models

type Notification struct {
	BaseModel
	User             User
	UserID           uint64 `sql:"type:bigint REFERENCES users(id)"`
	BookmarksShare   BookmarksShare
	BookmarksShareID uint64 `sql:"type:bigint REFERENCES bookmarks_shares(id)"`
}
