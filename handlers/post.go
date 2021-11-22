package handlers

import (
	"net/http"

	"github.com/bennysiahaan/twitter-clone/data"
)

// swagger:route POST /create tweets PostTweet
// Post a new tweet
//
// responses:
//  204: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Post handles POST requests to post new tweets to the database
func (t *Tweet) Post(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Tweet")

	Tweet := r.Context().Value(KeyTweet{}).(*data.Tweet)

    t.l.Printf("[DEBUG] Inserting tweet: %#v\n", Tweet)
    err := data.AddTweet(*Tweet)
	if err == data.ErrTweetNotFound {
		t.l.Println("[ERROR] tweet not found", err)
        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: "Tweet not found in database"}, rw)
		return
	}
	if err != nil {
        t.l.Println("[ERROR] creating tweet", err)
        rw.WriteHeader(http.StatusUnprocessableEntity)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
	}

    rw.WriteHeader(http.StatusNoContent)
}
