package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserMissingEmailReturnsErr(t *testing.T) {
	user := User{
		Username:    "test",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestUserMissingUsernameReturnsErr(t *testing.T) {
	user := User{
		Email:       "test@test.com",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestUserMissingDisplayNameReturnsErr(t *testing.T) {
	user := User{
		Email:    "test@test.com",
		Username: "test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidEmailReturnsErr1(t *testing.T) {
	user := User{
		Email:       "test@",
		Username:    "test",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidEmailReturnsErr2(t *testing.T) {
	user := User{
		Email:       "@test.com",
		Username:    "test",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidEmailReturnsErr3(t *testing.T) {
	user := User{
		Email:       "test@.com",
		Username:    "test",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidUsernameReturnsErr1(t *testing.T) {
	user := User{
		Email:       "test@test.com",
		Username:    "aa",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidUsernameReturnsErr2(t *testing.T) {
	user := User{
		Email:       "test@test.com",
		Username:    "aaaaaaaaaaaaaaaa",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidUsernameReturnsErr3(t *testing.T) {
	user := User{
		Email:       "test@test.com",
		Username:    "twitter",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidUsernameReturnsErr4(t *testing.T) {
	user := User{
		Email:       "test@test.com",
		Username:    "admin",
		DisplayName: "Test",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidDisplayNameReturnsErr1(t *testing.T) {
	user := User{
		Email:       "test@test.com",
		Username:    "test",
		DisplayName: "",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidDisplayNameReturnsErr2(t *testing.T) {
	user := User{
		Email:       "test@test.com",
		Username:    "test",
		DisplayName: "aaaaaaaaaaaaaaaaaaaaa",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestTweetInvalidProfileImageURLReturnsErr(t *testing.T) {
	user := User{
		Email:           "test@test.com",
		Username:        "test",
		DisplayName:     "Test",
		ProfileImageURL: "https://www.test.com/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 1)
}

func TestValidUserDoesNOTReturnErr(t *testing.T) {
	user := User{
		Email:           "test@test.com",
		Username:        "test",
		DisplayName:     "Test",
		ProfileImageURL: "https://www.test.com/test.jpg",
	}

	v := NewUserValidation()
	err := v.Validate(user)
	assert.Len(t, err, 0)
}

func TestUsersToJSON(t *testing.T) {
	users := []*User{
		{
			Email:           "test@test.com",
			Username:        "test",
			DisplayName:     "Test",
			ProfileImageURL: "https://www.test.com/test.jpg",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(users, b)
	assert.NoError(t, err)
}
