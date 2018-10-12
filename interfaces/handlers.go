package interfaces

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hlts2/gin-server-template/application"
	"github.com/hlts2/gin-server-template/domain"
	"github.com/jinzhu/gorm"
)

// GetTweet is handler to get tweet
func GetTweet(ctxt *gin.Context) {
	id := ctxt.GetString("id")
	if len(id) == 0 {
		ctxt.JSON(http.StatusBadRequest, errors.New("id is not full"))
		return
	}

	var tweet domain.Tweet

	var err error
	if tweet, err = application.GetTweet(id); err != nil {
		errorResponse(ctxt, err)
		return
	}

	ctxt.JSON(200, tweet)
}

// GetTweets is handler to get tweets
func GetTweets(ctxt *gin.Context) {
	var tweets domain.Tweets

	var err error
	if tweets, err = application.GetTweets(); err != nil {
		errorResponse(ctxt, err)
		return
	}

	ctxt.JSON(200, tweets)
}

func errorResponse(ctxt *gin.Context, err error) {
	if err == gorm.ErrRecordNotFound {
		ctxt.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
