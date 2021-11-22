package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bennysiahaan/twitter-clone/data"
	"github.com/gorilla/mux"
)

// KeyTweet is a key used for the Tweet object in the context
type KeyTweet struct{}

// Tweet handler for getting and updating tweets
type Tweet struct {
	l *log.Logger
	v *data.Validation
}

// NewTweet returns a new tweet handler with the given logger
func NewTweet(l *log.Logger, v *data.Validation) *Tweet {
	return &Tweet{l, v}
}

// ErrInvalidTweetPath is an error message when the tweet path is not valid
var ErrInvalidTweetPath = fmt.Errorf("Invalid Path, path should be /[tweetId]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// GetTweetIDParam returns the tweetId param from the URL
// Return the extracted param from the URL through regexp
// from the handler
func GetTweetIDParam(r *http.Request) string {
	// parse the tweet id from the url
	vars := mux.Vars(r)
	id := vars["tweetId"]

	return id
}
