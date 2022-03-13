package services

import (
	"encoding/json"
	"net/http"

	"github.com/obrkn/twitter/models"
	"github.com/obrkn/twitter/repositories"
	"github.com/obrkn/twitter/utils/logic"
	"github.com/obrkn/twitter/utils/validation"
)

type TweetService interface {
	GetAllTweets(w http.ResponseWriter, r *http.Request) ([]models.TweetResponse, error)
	CreateTweet(w http.ResponseWriter, r *http.Request) (models.TweetResponse, error)
	SendAllTweetsResponse(w http.ResponseWriter, tweets *[]models.TweetResponse)
	SendTweetResponse(w http.ResponseWriter, tweet *models.TweetResponse)
}

type tweetService struct {
	tr repositories.TweetRepository
	rl logic.ResponseLogic
	tv validation.TweetValidation
}

func NewTweetService(tr repositories.TweetRepository, rl logic.ResponseLogic, tv validation.TweetValidation) TweetService {
	return &tweetService{tr, rl, tv}
}

/*
	一覧取得
*/
func (ts *tweetService) GetAllTweets(w http.ResponseWriter, r *http.Request) ([]models.TweetResponse, error) {
	var tweets []models.TweetResponse

	err := ts.tr.GetAllTweets(&tweets)
	if err != nil {
		return []models.TweetResponse{}, err
	}

	return tweets, nil
}

/*
	投稿処理
*/
func (ts *tweetService) CreateTweet(w http.ResponseWriter, r *http.Request) (models.TweetResponse, error) {
	return models.TweetResponse{}, nil
}

/*
	Tweetリスト送信処理
*/
func (ts *tweetService) SendAllTweetsResponse(w http.ResponseWriter, tweets *[]models.TweetResponse) {
	response := *tweets
	responseBody, _ := json.Marshal(response)

	ts.rl.SendResponse(w, responseBody, http.StatusOK)
}

/*
	Tweet送信処理
*/
func (ts *tweetService) SendTweetResponse(w http.ResponseWriter, tweet *models.TweetResponse) {
}
