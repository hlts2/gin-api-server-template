package repository

import "github.com/hlts2/gin-server-template/domain"

// TweetRepository represetns repository interface of tweet
type TweetRepository interface {
	Get(id int) *domain.Tweet
	Gets() *domain.Tweets
}
