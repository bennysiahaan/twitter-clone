// Package classification Twitter Clone API
//
// Documentation for Twitter Clone API
//
//      Schemes: http
//      BasePath: /
//      Version: 1.0.0
//
//      Consumes:
//      - application/json
//
//      Produces:
//      - application/json
//
// swagger:meta
package handlers

import "github.com/bennysiahaan/twitter-clone/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
    // Description of the error
    // in: body
    Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
    // Collection of the errors
    // in: body
    Body ValidationError
}

// A list of tweets returns in the reponse
// swagger:response tweetsResponse
type tweetsResponseWrapper struct {
	// All tweets in the system
	// in: body
	Body []data.Tweet
}

// Data structure representing a single tweet
// swagger:response tweetResponse
type tweetResponseWrapper struct {
    // Newly created tweet
    // in: body
    Body data.Tweet
}

// No content is required by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters GetTweet DeleteTweet
type tweetIDParamsWrapper struct {
	// The id of the tweet for which the operation relates
	// in: path
	// required: true
	// length: 36
    // pattern: [0-9a-zA-Z-]{36}
	TweetID string `json:"tweetId"`
}
