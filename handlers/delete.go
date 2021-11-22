package handlers

import (
	"net/http"

	"github.com/bennysiahaan/twitter-clone/data"
)

// swagger:route DELETE /tweet/{tweetId} tweets DeleteTweet
// Delete a tweet
//
// responses:
//  204: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Delete handles DELETE requests to delete tweets from the database
func (t *Tweet) Delete(rw http.ResponseWriter, r *http.Request) {
    t.l.Println("Handle DELETE Tweets")

    id := GetTweetIDParam(r)

    t.l.Println("[DEBUG] Deleting record id", id)
    err := data.DeleteTweet(id)
	if err == data.ErrTweetNotFound {
		t.l.Println("[ERROR] tweet not found", err)
        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: "Tweet not found in database"}, rw)
		return
	}
	if err != nil {
        t.l.Println("[ERROR] deleting tweet", err)
        rw.WriteHeader(http.StatusUnprocessableEntity)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
	}

    rw.WriteHeader(http.StatusNoContent)
}