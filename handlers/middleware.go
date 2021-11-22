package handlers

import (
	"context"
	"net/http"

	"github.com/bennysiahaan/twitter-clone/data"
)

// MiddlewareValidateTweet validates the tweet in the request and calls next if ok
func (t *Tweet) MiddlewareValidateTweet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		Tweet := &data.Tweet{}

		err := data.FromJSON(Tweet, r.Body)
		if err != nil {
			t.l.Println("[ERROR] deserializing tweet", err)
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		errs := t.v.Validate(Tweet)
		if len(errs) != 0 {
			t.l.Println("[ERROR] validating tweet", err)
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		ctx := context.WithValue(r.Context(), KeyTweet{}, Tweet)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
