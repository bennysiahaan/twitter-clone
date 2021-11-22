package data

import (
	"database/sql"
	"fmt"

	"github.com/bennysiahaan/twitter-clone/db"
)

// ErrUserNotFound is an error raised when a tweet can not be found in the database
var ErrUserNotFound = fmt.Errorf("user not found")

// User defines the structure for an API user
// swagger:model
type User struct {
	UserID          string `json:"userId" validate:"omitempty"`
	Email           string `json:"email" validate:"required,user-email"`
	Username        string `json:"username" validate:"required,username"`
	DisplayName     string `json:"displayName" validate:"required,user-displayName"`
	ProfileImageURL string `json:"avatarUrl" validate:"omitempty,user-avatarUrl"`
}

// UserInfo defines the structure for user information when fetching
// a tweet from this user
type UserInfo struct {
	Username        string
	DisplayName     string
	ProfileImageURL string
}

// GetUserInfo returns user info for a given tweet's userId
// If the user info could not be retrived from the database
// this function returns a DB error
func GetUserInfo(id string) (*UserInfo, error) {
	DB := db.GetDB()
	var ui UserInfo

	row := DB.QueryRow("SELECT username, display_name, profile_image_url FROM users WHERE user_id = ?", id)
	if err := row.Scan(
		&ui.Username,
		&ui.DisplayName,
		&ui.ProfileImageURL,
	); err != nil {
		if err == sql.ErrNoRows {
			return &ui, ErrUserNotFound
		}
		return &ui, err
	}

	return &ui, nil
}
