package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hlts2/gin-server-template/domain"
	"github.com/hlts2/gin-server-template/domain/repository"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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
		if err == gorm.ErrRecordNotFound {
			return tweet, errors.Wrap(err, "tweet id#"+id)
		}
		return tweet, err
	}

	return tweet, nil
}

func (tr *tweetRepositoryImpl) Gets() (domain.Tweets, error) {
	tweets := domain.Tweets{}
	if err := tr.conn.Find(&tweets).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return tweets, errors.Wrap(err, "tweets ")
		}
		return tweets, err
	}

	if len(tweets) == 0 {
		return tweets, errors.Wrap(gorm.ErrRecordNotFound, "tweets ")
	}

	return tweets, nil
}
