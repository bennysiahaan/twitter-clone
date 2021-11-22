package handlers

import (
	"net/http"

	"github.com/bennysiahaan/twitter-clone/data"
)

// swagger:route GET / tweets GetTimeline
// Return the user's timeline from the database
// responses:
//  200: tweetsResponse
//  404: errorResponse
//  500: errorResponse

// GetTimeline handles GET requests and returns all tweets from a user's timeline
func (t *Tweet) GetTimeline(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Timeline")

	rw.Header().Add("Content-Type", "application/json")

	Tweets, err := data.GetTweets()
	if err != nil {
		t.l.Println("[ERROR] fetching records", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	type TimelineItem struct {
		TweetID     string `json:"tweetId"`
		UserID      string `json:"userId"`
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
		AvatarURL   string `json:"avatarUrl"`
		Body        string `json:"body"`
		ContentURL  string `json:"contentUrl"`
		UpdatedAt   string `json:"updatedAt"`
	}

	var Timeline []*TimelineItem

	for _, tweet := range Tweets {
		uid := tweet.UserID
		ui, err := data.GetUserInfo(uid)
		if err != nil {
			t.l.Println("[ERROR] fetching records", err)
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		Timeline = append(Timeline, &TimelineItem{
			TweetID:     tweet.TweetID,
			UserID:      tweet.UserID,
			Username:    ui.Username,
			DisplayName: ui.DisplayName,
			AvatarURL:   ui.ProfileImageURL,
            Body: tweet.Body,
            ContentURL: tweet.ContentURL,
            UpdatedAt: tweet.UpdatedAt,
		})
	}

	err = data.ToJSON(Timeline, rw)
	if err != nil {
		t.l.Println("[ERROR] serializing records", err)
		return
	}
}

// swagger:route GET /tweet/{tweetId} tweets GetTweet
// Return a tweet that matches the given tweetId from the database
// responses:
//  200: tweetResponse
//  404: errorResponse

// GetTweet handles GET requests
func (t *Tweet) GetTweet(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Tweet")

	id := GetTweetIDParam(r)

	t.l.Println("[DEBUG] get record id", id)

	tweet, err := data.GetTweetByID(id)

	switch err {
	case nil:
	case data.ErrTweetNotFound:
		t.l.Println("[ERROR] fetching tweet", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		t.l.Println("[ERROR] fetching tweet", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(tweet, rw)
	if err != nil {
		t.l.Println("[ERROR] serializing tweet", err)
	}
}
