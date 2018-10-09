package persistence

import (
	"github.com/hlts2/gin-server-template/domain"
	"github.com/hlts2/gin-server-template/domain/repository"
	"github.com/jinzhu/gorm"
)

// tweetRepositoryImpl is implementation of TweetRepository interfaces
type tweetRepositoryImpl struct {
	conn *gorm.DB
}

// NewTweetRepositoryImplWithRDB returns tweetRepositoryImpl(repository.TweetRepository) object
func NewTweetRepositoryImplWithRDB(conn *gorm.DB) repository.TweetRepository {
	return &tweetRepositoryImpl{conn: conn}
}

func (tr *tweetRepositoryImpl) Get(id string) *domain.Tweet {
	tweet := &domain.Tweet{}
	tweet.ID = id
	tr.conn.First(tweet)
	return tweet
}

func (tr *tweetRepositoryImpl) Gets() domain.Tweets {
	tweets := domain.Tweets{}
	tr.conn.Find(&tweets)
	return tweets
}
