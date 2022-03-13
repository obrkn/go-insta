package repositories

import (
	"database/sql"

	"github.com/obrkn/twitter/models"
)

type TweetRepository interface {
	GetAllTweets(tweets *[]models.TweetResponse) error
	CreateTweet(createTweet *models.Tweet, userId int) error
}

type tweetRepository struct {
	db *sql.DB
}

func NewTweetRepository(db *sql.DB) TweetRepository {
	return &tweetRepository{db}
}

/*
	全件ツイート取得
*/
func (tr *tweetRepository) GetAllTweets(tweets *[]models.TweetResponse) error {
	// rows, err := tr.db.Query("SELECT id, message, created_at, updated_at, user FROM tweets")
	rows, err := tr.db.Query("SELECT T.id, T.message, T.created_at, T.updated_at, U.nickname FROM tweets T JOIN users U ON U.id = T.user_id")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tweet models.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.Message, &tweet.CreatedAt, &tweet.UpdatedAt, &tweet.UserNickname); err != nil {
			return err
		}
		*tweets = append(*tweets, tweet)
	}

	return nil
}

/*
	新規ツイート登録
*/
func (tr *tweetRepository) CreateTweet(createTweet *models.Tweet, userId int) error {
	_, err := tr.db.Exec("INSERT INTO tweets(user_id, message) VALUES(?, ?);", userId, createTweet.Message)
	if err != nil {
		return err
	}

	return nil
}
