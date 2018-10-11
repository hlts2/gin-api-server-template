package persistence

import (
	_ "github.com/go-sql-driver/mysql"
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

func (tr *tweetRepositoryImpl) Get(id string) (domain.Tweet, error) {
	tweet := domain.Tweet{}
	tweet.ID = id

	if err := tr.conn.First(&tweet).Error; err != nil {
		return tweet, err
	}

	return tweet, nil
}

func (tr *tweetRepositoryImpl) Gets() (domain.Tweets, error) {
	tweets := domain.Tweets{}
	if err := tr.conn.Find(&tweets).Error; err != nil {
		return tweets, err
	}

	return tweets, nil
}
