package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/obrkn/twitter/controllers"
)

type TweetRouter interface {
	SetTweetRouting(api *mux.Router)
}

type tweetRouter struct {
	tc controllers.TweetController
}

func NewTweetRouter(tc controllers.TweetController) TweetRouter {
	return &tweetRouter{tc}
}

func (tr *tweetRouter) SetTweetRouting(api *mux.Router) {
	api.HandleFunc("/tweets", tr.tc.FetchAllTweets).Methods(http.MethodGet)
	api.HandleFunc("/tweets", tr.tc.CreateTweet).Methods(http.MethodPost, http.MethodOptions)
}
