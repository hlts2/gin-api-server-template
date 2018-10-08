package persistence

import (
	"github.com/hlts2/gin-server-template/domain"
	"github.com/hlts2/gin-server-template/domain/repository"
)

// tweetRepositoryImpl is implementation of TweetRepository interfaces
type tweetRepositoryImpl struct{}

// NewTweetRepositoryImplWithRDB returns tweetRepositoryImpl(repository.TweetRepository) object
func NewTweetRepositoryImplWithRDB() repository.TweetRepository {
	return &tweetRepositoryImpl{}
}

func (tr *tweetRepositoryImpl) Get(id int) *domain.Tweet {
	return nil
}

func (tr *tweetRepositoryImpl) Gets() *domain.Tweets {
	return nil
}
