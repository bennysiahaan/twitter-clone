package handlers

import (
	"net/http"
)

// RedirectHome redirects to /home
// Panics if cannot find the handler for /home
func (t *Tweet) RedirectHome(rw http.ResponseWriter, r *http.Request) {
	url := "http://" + r.Host + "/home"
	http.Redirect(rw, r, url, http.StatusMovedPermanently)
}

// RedirectTweet redirects /{tweetId} to /tweet/{tweetId}
// Panics if cannot find the handler for /tweet/{tweetId}
func (t *Tweet) RedirectTweet(rw http.ResponseWriter, r *http.Request) {
	url := "http://" + r.Host + "/tweet" + r.URL.String()
	http.Redirect(rw, r, url, http.StatusFound)
}
