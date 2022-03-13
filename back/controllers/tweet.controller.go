package controllers

import (
	"net/http"

	"github.com/obrkn/twitter/services"
)

type TweetController interface {
	FetchAllTweets(w http.ResponseWriter, r *http.Request)
	CreateTweet(w http.ResponseWriter, r *http.Request)
}

type tweetController struct {
	ts services.TweetService
	as services.AuthService
}

func NewTweetController(ts services.TweetService, as services.AuthService) TweetController {
	return &tweetController{ts, as}
}

/*
 一覧取得
*/
func (tc *tweetController) FetchAllTweets(w http.ResponseWriter, r *http.Request) {
	// 一覧取得
	tweets, err := tc.ts.GetAllTweets(w, r)
	if err != nil {
		return
	}

	// レスポンス送信
	tc.ts.SendAllTweetsResponse(w, &tweets)
}

/*
	投稿
*/
func (tc *tweetController) CreateTweet(w http.ResponseWriter, r *http.Request) {
	// プリフライト
	if r.Method == http.MethodOptions {
		tc.as.Preflight(w, r)
		return
	}

	// ログイン
	user, err := tc.ts.CreateTweet(w, r)
	if err != nil {
		return
	}

	// レスポンス送信
	tc.ts.SendTweetResponse(w, &user)
}
