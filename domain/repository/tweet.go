package repository

import "github.com/hlts2/gin-server-template/domain"

// TweetRepository represetns repository interface of tweet
type TweetRepository interface {
	Get(id string) *domain.Tweet
	Gets() domain.Tweets
}
