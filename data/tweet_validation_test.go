package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTweetMissingUserIDReturnsErr(t *testing.T) {
    tweet := Tweet{
        Body: "test",
    }

    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 1)
}

func TestTweetMissingBodyReturnsErr(t *testing.T) {
    tweet := Tweet{
        UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
    }

    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 1)
}

func TestTweetInvalidIDReturnsErr1(t *testing.T) {
    tweet := Tweet{
        TweetID: "test",
        UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
        Body: "test",
    }
    
    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 1)
}

func TestTweetInvalidIDReturnsErr2(t *testing.T) {
    tweet := Tweet{
        TweetID: "test-52765eff-4470-11ec-b65e-00155d0a220d",
        UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
        Body: "test",
    }
    
    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 1)
}

func TestTweetInvalidUserIDReturnsErr(t *testing.T) {
    tweet := Tweet{
        UserID: "test",
        Body: "test",
    }
    
    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 1)
}

func TestTweetInvalidBodyReturnsErr1(t *testing.T) {
    tweet := Tweet{
        UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
        Body: "",
    }

    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 1)
}

func TestTweetInvalidBodyReturnsErr2(t *testing.T) {
    tweet := Tweet{
        UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
        Body: "asfasdfasdfdsafsdfadssadfasdfasdfsadfasdfasffawe qwefybvqweku vqwk uqgv kfuqwgev kfjqwgevf kqwjeh vkjwq vkqjwhvqkwjh fvkjqw fkjqwhfkjwh vkqjwhvfkjqwhevkqjhevkjhqwevkhwqev kjhqvwekhvwqekvhw qvkqjwh vkjhwev fkjhqvwe jvewkjq vkjqwhvkqwve kjvwehkvew hvqwek hvqwe khve k qejhfv qwejhvas",
    }

    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 1)
}

func TestValidTweetDoesNOTReturnErr1(t *testing.T) {
    tweet := Tweet{
        UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
        Body: "test",
    }
    
    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 0)
}

func TestValidTweetDoesNOTReturnErr2(t *testing.T) {
    tweet := Tweet{
        TweetID: "52765eff-4470-11ec-b65e-00155d0a220d",
        UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
        Body: "test",
    }
    
    v := NewTweetValidation()
    err := v.Validate(tweet)
    assert.Len(t, err, 0)
}

func TestTweetsToJSON1(t *testing.T) {
    tweets := []*Tweet{
        {
            UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
            Body: "test",
        },
    }
    
    b := bytes.NewBufferString("")
    err := ToJSON(tweets, b)
    assert.NoError(t, err)
}

func TestTweetsToJSON2(t *testing.T) {
    tweets := []*Tweet{
        {
            TweetID: "52765eff-4470-11ec-b65e-00155d0a220d",
            UserID: "52765eff-4470-11ec-b65e-00155d0a220d",
            Body: "test",
        },
    }
    
    b := bytes.NewBufferString("")
    err := ToJSON(tweets, b)
    assert.NoError(t, err)
}