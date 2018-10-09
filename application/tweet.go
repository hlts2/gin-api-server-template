package application

import (
	"github.com/hlts2/gin-server-template/domain"
	"github.com/hlts2/gin-server-template/infrastructure"
	"github.com/hlts2/gin-server-template/infrastructure/persistence"
)

// GetTweet returns tweet model object of given id
func GetTweet(id string) *domain.Tweet {
	conn, err := infrastructure.NewDBConnection()
	if err != nil {
		return nil
	}
	defer conn.Close()

	repo := persistence.NewTweetRepositoryImplWithRDB(conn)
	return repo.Get(id)
}

// GetTweets returns all tweet models
func GetTweets() domain.Tweets {
	conn, err := infrastructure.NewDBConnection()
	if err != nil {
		return nil
	}

	defer conn.Close()

	rep := persistence.NewTweetRepositoryImplWithRDB(conn)
	return rep.Gets()
}
