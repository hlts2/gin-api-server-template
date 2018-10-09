package persistence

import (
	"database/sql"

	"github.com/hlts2/gin-server-template/domain"
	"github.com/hlts2/gin-server-template/domain/repository"
)

// tweetRepositoryImpl is implementation of TweetRepository interfaces
type tweetRepositoryImpl struct {
	conn *sql.DB
}

// NewTweetRepositoryImplWithRDB returns tweetRepositoryImpl(repository.TweetRepository) object
func NewTweetRepositoryImplWithRDB(conn *sql.DB) repository.TweetRepository {
	return &tweetRepositoryImpl{conn: conn}
}

func (tr *tweetRepositoryImpl) Get(id string) *domain.Tweet {
	return nil
}

func (tr *tweetRepositoryImpl) Gets() *domain.Tweets {
	return nil
}
