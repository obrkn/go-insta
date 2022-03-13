package validation

type TweetValidation interface {
}

type tweetValidation struct{}

func NewTweetValidation() TweetValidation {
	return &tweetValidation{}
}
