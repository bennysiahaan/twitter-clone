package data

import (
	"database/sql"
	"fmt"

	"github.com/bennysiahaan/twitter-clone/db"
)

// ErrTweetNotFound is an error raised when a tweet can not be found in the database
var ErrTweetNotFound = fmt.Errorf("tweet not found")

// Tweet defines the structure for an API tweet
// swagger:model
type Tweet struct {
	// the id for the tweet (auto-generated)
	//
	// required: false
	// pattern: [0-9a-zA-Z-]{36}
	TweetID string `json:"tweetId" validate:"omitempty,tweet-id"` // Unique identifier for the product

	// the id of the user that posts this tweet
	//
	// required: true
	// pattern: [0-9a-zA-Z-]{36}
	UserID string `json:"userId" validate:"required,tweet-userId"`

	// the body of this tweet
	//
	// required: true
	// min length: 1
	// max length: 280
	Body string `json:"body" validate:"required,tweet-body"`

	// the URL of embedded content
	//
	// required: false
	// max length: 1024
	ContentURL string `json:"contentUrl"`

	// the date when this tweet was created
	//
	// required: false
	CreatedAt string `json:"createdAt"`

	// the date when this tweet was last modified
	//
	// required: false
	UpdatedAt string `json:"updatedAt"`
}

// Tweets defines a slice of Tweet
type Tweets []*Tweet

// GetTweets returns all tweets for the user's timeline
// If tweets could not be retrieved from the database
// this function returns a DB error
func GetTweets() (Tweets, error) {
	var tweets Tweets
	DB := db.GetDB()

	rows, err := DB.Query("SELECT tweet_id, user_id, body, content_url, created_at, updated_at FROM tweets ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tweet Tweet
		if err := rows.Scan(
			&tweet.TweetID,
			&tweet.UserID,
			&tweet.Body,
			&tweet.ContentURL,
			&tweet.CreatedAt,
			&tweet.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tweets = append(tweets, &tweet)
	}

	return tweets, nil
}

// GetTweetByID returns a single tweet which matches the tweet_id from the database
// If a tweet is not found this function returns a TweetNotFound error
// Otherwise if the tweet is found but could not be retrieved from the database
// this function returns a DB error
func GetTweetByID(id string) (*Tweet, error) {
	DB := db.GetDB()
	var tweet Tweet

	row := DB.QueryRow("SELECT tweet_id, user_id, body, content_url, created_at, updated_at FROM tweets WHERE tweet_id = ?", id)
	if err := row.Scan(
		&tweet.TweetID,
		&tweet.UserID,
		&tweet.Body,
		&tweet.ContentURL,
		&tweet.CreatedAt,
		&tweet.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return &tweet, ErrTweetNotFound
		}
		return &tweet, err
	}

	return &tweet, nil
}

// AddTweet adds a new tweet to the database
// This function returns a DB error if the tweet could not be added to the database
func AddTweet(t Tweet) error {
	DB := db.GetDB()
	if err := DB.Ping(); err != nil {
		return err
	}

	result, err := DB.Exec(
		"INSERT INTO tweets (tweet_id, user_id, body, content_url) VALUES (uuid(), ?, ?, ?)",
		t.UserID,
		t.Body,
		t.ContentURL,
	)
	if n, err := result.RowsAffected(); err == nil {
		if n == 0 {
			return ErrTweetNotFound
		}
	}

	return err
}

// EditTweet changes the body of a tweet that matches the given id
// If a tweet with the given id does not exist in the database
// or if it could not be edited in the database
// this function returns a TweetNotFound error or a DB error respectively
func EditTweet(t Tweet) error {
	DB := db.GetDB()

	result, err := DB.Exec("UPDATE tweets SET body = ? WHERE tweet_id = ?", t.Body, t.TweetID)
	if n, err := result.RowsAffected(); err == nil {
		if n == 0 {
			return ErrTweetNotFound
		}
	}

	return err
}

// DeleteTweet deletes a tweet that matches the given id
// If a tweet with the given id does not exist in the database
// or if it could not be edited in the database
// this function returns a TweetNotFound error or a DB error respectively
func DeleteTweet(id string) error {
	DB := db.GetDB()

	result, err := DB.Exec("DELETE FROM tweets WHERE tweet_id = ?", id)
	if n, err := result.RowsAffected(); err == nil {
		if n == 0 {
			return ErrTweetNotFound
		}
	}

	return err
}
