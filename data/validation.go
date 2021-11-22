package data

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/go-playground/validator"
)

type ValidationError struct {
    validator.FieldError
}

func (v ValidationError) Error() string {
    return fmt.Sprintf(
        "Key: '%s' Error: Field validation for '%s' failed on the '%s' tag\n",
        v.Namespace(),
        v.Field(),
        v.Tag(),
    )
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
    errs := []string{}
    for _, err := range v {
        errs = append(errs, err.Error())
    }

    return errs
}

// Validation contains
type Validation struct {
    validate *validator.Validate
}

// NewTweetValidation creates a new Validation type for tweets
func NewTweetValidation() *Validation {
    validate := validator.New()
    validate.RegisterValidation("tweet-id", ValidateTweetID)
    validate.RegisterValidation("tweet-userId", ValidateUserID)
    validate.RegisterValidation("tweet-body", ValidateBody)
    return &Validation{validate}
}

// NewUserValidation creates a new Validation type for users
func NewUserValidation() *Validation {
    validate := validator.New()
    validate.RegisterValidation("user-email", ValidateEmail)
    validate.RegisterValidation("username", ValidateUsername)
    validate.RegisterValidation("user-display-name", ValidateDisplayName)
    validate.RegisterValidation("user-profile-img-url", ValidateProfileImageURL)
    return &Validation{validate}
}

// Validate tweet
// for more details the returned error can be cast into a
// validator.ValidationErrors collection
//
// if ve, ok := err.(validator.ValidationErrors); ok {
//     fmt.Println(ve.Namespace())
//     fmt.Println(ve.Field())
//     fmt.Println(ve.StructNamespace())
//     fmt.Println(ve.StructField())
//     fmt.Println(ve.Tag())
//     fmt.Println(ve.ActualTag())
//     fmt.Println(ve.Kind())
//     fmt.Println(ve.Type())
//     fmt.Println(ve.Value())
//     fmt.Println(ve.Param())
//     fmt.Println()
// }
func (v *Validation) Validate(i interface{}) ValidationErrors {
    errs, ok := v.validate.Struct(i).(validator.ValidationErrors)

    if(!ok) {
        return nil
    }

    if len(errs) == 0 {
        return nil
    }

    var returnErrs []ValidationError
    for _, err := range errs {
        ve := ValidationError{err}
        returnErrs = append(returnErrs, ve)
    }

    return returnErrs
}

// Validate the id of the tweet
func ValidateTweetID(fl validator.FieldLevel) bool {
    return len(fl.Field().String()) == 36
}

// Validate the userId of the tweet
func ValidateUserID(fl validator.FieldLevel) bool {
    return len(fl.Field().String()) == 36
}

// Validate the body of the tweet
func ValidateBody(fl validator.FieldLevel) bool {
    return len(fl.Field().String()) >= 0 && len(fl.Field().String()) <= 280
}

// Validate the user's email
func ValidateEmail(fl validator.FieldLevel) bool {
    _, err := mail.ParseAddress(fl.Field().String())
    return err == nil
}

// Validate the username
func ValidateUsername(fl validator.FieldLevel) bool {
    re1 := regexp.MustCompile(`^(\w{3,15})$`)
    re2 := regexp.MustCompile("(twitter)")
    re3 := regexp.MustCompile("(admin)")
    un := fl.Field().String()
    return re1.MatchString(un) && !re2.MatchString(un) && !re3.MatchString(un)
}

// Validate display name
func ValidateDisplayName(fl validator.FieldLevel) bool {
    dn := fl.Field().String()
    return len(dn) >= 1 && len(dn) <= 20
}

// Validate profile image URL
func ValidateProfileImageURL(fl validator.FieldLevel) bool {
    url := fl.Field().String()
    return len(url) <= 1024
}