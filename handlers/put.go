package handlers

import (
	"net/http"

	"github.com/bennysiahaan/twitter-clone/data"
)

// swagger:route PUT /edit tweets EditTweet
// Edit the body of a tweet
//
// responses:
//  204: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Edit handles PUT requests to edit tweets in the database
func (t *Tweet) Edit(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PUT Tweet")

	Tweet := r.Context().Value(KeyTweet{}).(*data.Tweet)

	t.l.Println("[DEBUG] Updating record id", Tweet.TweetID)
	err := data.EditTweet(*Tweet)
	if err == data.ErrTweetNotFound {
		t.l.Println("[ERROR] tweet not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Tweet not found in database"}, rw)
		return
	}
	if err != nil {
		t.l.Println("[ERROR] updating tweet", err)
		rw.WriteHeader(http.StatusUnprocessableEntity)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
